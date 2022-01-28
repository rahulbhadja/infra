package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/infrahq/infra/internal/api"
	"github.com/infrahq/infra/internal/logging"
	"github.com/infrahq/infra/uid"
)

type UseOptions struct {
	Name      string
	Namespace string
	Labels    []string `mapstructure:"labels"`
}

func use(options *UseOptions) error {
	config, err := currentHostConfig()
	if err != nil {
		return err
	}

	client, err := defaultAPIClient()
	if err != nil {
		return err
	}

	users, err := client.ListUsers(config.Name)
	if err != nil {
		if errors.Is(err, api.ErrForbidden) {
			fmt.Fprintln(os.Stderr, "Session has expired.")

			if err = login(&LoginOptions{Current: true}); err != nil {
				return err
			}

			return use(options)
		}

		return err
	}

	// This shouldn't be possible but check nonetheless
	switch {
	case len(users) < 1:
		//lint:ignore ST1005, user facing error
		return fmt.Errorf("User %q not found", config.Name)
	case len(users) > 1:
		//lint:ignore ST1005, user facing error
		return fmt.Errorf("Found multiple users for %q, please contact your administrator", config.Name)
	}

	user := users[0]

	// first make sure kubeconfig is up to date
	if err := updateKubeconfig(user); err != nil {
		return err
	}

	// deduplicate candidates
	candidates := make(map[string][]api.Grant)
	for _, r := range user.Grants {
		candidates[r.Destination.NodeID] = append(candidates[r.Destination.NodeID], r)
	}

	for _, g := range user.Groups {
		for _, r := range g.Grants {
			candidates[r.Destination.NodeID] = append(candidates[r.Destination.NodeID], r)
		}
	}

	// find candidate destinations
	destinations := make(map[string]map[string][]api.Grant)

DESTINATIONS:
	for _, d := range candidates {
		for _, r := range d {
			logging.S.Debugf("considering %s %s@%s#%s", r.ID, r.Destination.Name, r.Destination.NodeID[:12], r.Kubernetes.Namespace)
			switch options.Name {
			case "":
			case r.Destination.Name:
			case r.Destination.NodeID:
			case r.Destination.NodeID[:12]:
			default:
				continue
			}

			switch options.Namespace {
			case "":
			case r.Kubernetes.Namespace:
			default:
				continue
			}

			labels := make(map[string]bool)
			for _, l := range r.Destination.Labels {
				labels[l] = true
			}

			for _, l := range options.Labels {
				if _, ok := labels[l]; !ok {
					continue DESTINATIONS
				}
			}

			if _, ok := destinations[r.Destination.NodeID[:12]]; !ok {
				destinations[r.Destination.NodeID[:12]] = make(map[string][]api.Grant)
			}

			destinations[r.Destination.NodeID[:12]][r.Kubernetes.Namespace] = append(destinations[r.Destination.NodeID[:12]][r.Kubernetes.Namespace], r)
		}
	}

	logging.S.Debugf("found %d suitable destination(s)", len(destinations))

	var namespaces map[string][]api.Grant

	switch len(destinations) {
	case 0:
		//lint:ignore ST1005, user facing error
		return fmt.Errorf("No kubernetes contexts found for user, you are not assigned any kubernetes grants")
	case 1:
		for _, d := range destinations {
			namespaces = d
		}
	default:
		promptOptions := make([]string, 0)

		for k, c := range destinations {
			// sample one namespace for this destinations
			var sample api.Grant
			for _, n := range c {
				sample = n[0]
				break
			}

			promptOptions = append(promptOptions, fmt.Sprintf("%s %s [%s]", k, sample.Destination.Name, strings.Join(sample.Destination.Labels, ", ")))
		}

		sort.Slice(promptOptions, func(i, j int) bool {
			return promptOptions[i] < promptOptions[j]
		})

		prompt := survey.Select{
			Message: "Select a cluster:",
			Options: promptOptions,
		}

		var selected string

		err := survey.AskOne(&prompt, &selected, survey.WithStdio(os.Stdin, os.Stderr, os.Stderr))
		if err != nil {
			if errors.Is(err, terminal.InterruptErr) {
				return nil
			}

			return err
		}

		parts := strings.Split(selected, " ")
		namespaces = destinations[parts[0]]
	}

	logging.S.Debugf("found %d suitable namespace(s)", len(namespaces))

	var namespace api.Grant

	switch len(namespaces) {
	case 0:
		// should be impossible
		//lint:ignore ST1005, user facing error
		return fmt.Errorf("No namespaces found for kubernetes contexts, your server configuration may be invalid")
	case 1:
		for _, n := range namespaces {
			namespace = n[0]
		}
	default:
		promptOptions := make([]string, 0)

		for _, n := range namespaces {
			names := make([]string, 0)

			var namespace string

			for _, r := range n {
				names = append(names, r.Kubernetes.Name)
				namespace = r.Kubernetes.Namespace
			}

			if namespace == "" {
				namespace = "*"
			}

			promptOptions = append(promptOptions, fmt.Sprintf("%s [%s]", namespace, strings.Join(names, ", ")))
		}

		sort.Slice(promptOptions, func(i, j int) bool {
			return promptOptions[i] < promptOptions[j]
		})

		prompt := survey.Select{
			Message: "Select a namespace:",
			Options: promptOptions,
		}

		var selected string

		err := survey.AskOne(&prompt, &selected, survey.WithStdio(os.Stdin, os.Stderr, os.Stderr))
		if err != nil {
			if errors.Is(err, terminal.InterruptErr) {
				return nil
			}

			return err
		}

		parts := strings.Split(selected, " ")
		if parts[0] == "*" {
			parts[0] = ""
		}

		namespace = namespaces[parts[0]][0]
	}

	if err := kubernetesSetContext(namespace.Destination.Name, namespace.Destination.NodeID[:12], namespace.Kubernetes.Namespace); err != nil {
		return err
	}

	return nil
}

func clientConfig() clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.WarnIfAllMissing = false

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})
}

func kubernetesSetContext(alias string, shortname string, namespace string) error {
	config := clientConfig()

	kubeconfig, err := config.RawConfig()
	if err != nil {
		return err
	}

	if c, ok := kubeconfig.Contexts[fmt.Sprintf("infra:%s@%s:%s", alias, shortname, namespace)]; ok {
		// try infra:<ALIAS>@<SHORTNAME>:<NAMESPACE>
		kubeconfig.CurrentContext = c.Cluster
	} else if c, ok := kubeconfig.Contexts[fmt.Sprintf("infra:%s:%s", alias, namespace)]; ok {
		// try infra:<ALIAS>:<NAMESPACE>
		kubeconfig.CurrentContext = c.Cluster
	} else if c, ok := kubeconfig.Contexts[fmt.Sprintf("infra:%s@%s", alias, shortname)]; ok {
		// try infra:<ALIAS>@<SHORTNAME>
		kubeconfig.CurrentContext = c.Cluster
	} else if c, ok := kubeconfig.Contexts[fmt.Sprintf("infra:%s", alias)]; ok {
		// try infra:<ALIAS>
		kubeconfig.CurrentContext = c.Cluster
	} else {
		return fmt.Errorf("Infra context not found in local Kubernetes configuration, Infra context should be created on login")
	}

	fmt.Fprintf(os.Stderr, "Switched to context %q.\n", kubeconfig.CurrentContext)

	if err := clientcmd.WriteToFile(kubeconfig, config.ConfigAccess().GetDefaultFilename()); err != nil {
		return err
	}

	return nil
}

func updateKubeconfig(user api.User) error {
	defaultConfig := clientConfig()

	kubeConfig, err := defaultConfig.RawConfig()
	if err != nil {
		return err
	}

	// deduplicate grants
	aliases := make(map[string]map[string]bool)
	grants := make(map[uid.ID]api.Grant)

	for _, r := range user.Grants {
		if _, ok := aliases[r.Destination.Name]; !ok {
			aliases[r.Destination.Name] = make(map[string]bool)
		}

		aliases[r.Destination.Name][r.Destination.NodeID] = true
		grants[r.ID] = r
	}

	for _, g := range user.Groups {
		for _, r := range g.Grants {
			if _, ok := aliases[r.Destination.Name]; !ok {
				aliases[r.Destination.Name] = make(map[string]bool)
			}

			aliases[r.Destination.Name][r.Destination.NodeID] = true
			grants[r.ID] = r
		}
	}

	for _, grant := range grants {
		name := grant.Destination.NodeID[:12]
		alias := grant.Destination.Name

		// TODO (#546): allow user to specify prefix, default ""
		// format: "infra:<ALIAS>"
		contextName := fmt.Sprintf("infra:%s", alias)

		if len(aliases[alias]) > 1 {
			// disambiguate destination by appending the ID
			// format: "infra:<ALIAS>@<NAME>"
			contextName = fmt.Sprintf("%s@%s", contextName, name)
		}

		if grant.Kubernetes.Namespace != "" {
			// destination is scoped to a namespace
			// format: "infra:<ALIAS>[@<NAME>]:<NAMESPACE>"
			contextName = fmt.Sprintf("%s:%s", contextName, grant.Kubernetes.Namespace)
		}

		logging.S.Debugf("creating kubeconfig for %s", contextName)

		kubeConfig.Clusters[contextName] = &clientcmdapi.Cluster{
			Server:                   fmt.Sprintf("https://%s/proxy", grant.Destination.Kubernetes.Endpoint),
			CertificateAuthorityData: []byte(grant.Destination.Kubernetes.CA),
		}

		kubeConfig.Contexts[contextName] = &clientcmdapi.Context{
			Cluster:   contextName,
			AuthInfo:  contextName,
			Namespace: grant.Kubernetes.Namespace,
		}

		executable, err := os.Executable()
		if err != nil {
			return err
		}

		kubeConfig.AuthInfos[contextName] = &clientcmdapi.AuthInfo{
			Exec: &clientcmdapi.ExecConfig{
				Command:         executable,
				Args:            []string{"tokens", "create", grant.Destination.NodeID},
				APIVersion:      "client.authentication.k8s.io/v1beta1",
				InteractiveMode: clientcmdapi.IfAvailableExecInteractiveMode,
			},
		}
	}

	for contextName := range kubeConfig.Contexts {
		parts := strings.Split(contextName, ":")

		// shouldn't be possible but we don't actually care
		if len(parts) < 1 {
			continue
		}

		if parts[0] != "infra" {
			continue
		}

		found := false

		for _, r := range grants {
			parts := strings.Split(parts[1], "@")

			switch {
			case len(parts) == 1:
				found = parts[0] == r.Destination.Name
			case len(parts) > 1:
				found = parts[0] == r.Destination.Name && parts[1] == r.Destination.NodeID[:12]
			}

			if found {
				break
			}
		}

		if !found {
			delete(kubeConfig.Clusters, contextName)
			delete(kubeConfig.Contexts, contextName)
			delete(kubeConfig.AuthInfos, contextName)
		}
	}

	kubeConfigFilename := defaultConfig.ConfigAccess().GetDefaultFilename()

	if err := clientcmd.WriteToFile(kubeConfig, kubeConfigFilename); err != nil {
		return err
	}

	return nil
}

func switchToFirstInfraContext() (string, error) {
	defaultConfig := clientConfig()

	kubeConfig, err := defaultConfig.RawConfig()
	if err != nil {
		return "", err
	}

	resultContext := ""

	if kubeConfig.Contexts[kubeConfig.CurrentContext] != nil && strings.HasPrefix(kubeConfig.CurrentContext, "infra:") {
		// if the current context is an infra-controlled context, stay there
		resultContext = kubeConfig.CurrentContext
	} else {
		for _, c := range kubeConfig.Contexts {
			if !strings.HasPrefix(c.Cluster, "infra:") {
				continue
			}

			// prefer a context with "default" or no namespace
			if c.Namespace == "" || c.Namespace == "default" {
				resultContext = c.Cluster
				break
			}

			resultContext = c.Cluster
		}
	}

	if resultContext != "" {
		kubeConfig.CurrentContext = resultContext
		if err = clientcmd.WriteToFile(kubeConfig, defaultConfig.ConfigAccess().GetDefaultFilename()); err != nil {
			return "", err
		}
	}

	return resultContext, nil
}
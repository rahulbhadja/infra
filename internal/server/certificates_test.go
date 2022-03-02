package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/infrahq/infra/internal/server/models"
	"github.com/infrahq/infra/pki"
	"github.com/infrahq/infra/uid"
	"github.com/stretchr/testify/require"
)

func TestCertificateSigningWorks(t *testing.T) {
	db := setupDB(t)

	cp, err := pki.NewNativeCertificateProvider(db, pki.NativeCertificateProviderConfig{
		FullKeyRotationDurationInDays: 2,
	})
	require.NoError(t, err)

	err = cp.CreateCA()
	require.NoError(t, err)

	err = cp.RotateCA()
	require.NoError(t, err)

	user := &models.User{
		Model: models.Model{ID: uid.New()},
		Email: "joe@example.com",
	}

	keyPair, err := pki.MakeUserCert("User "+user.ID.String(), 24*time.Hour)
	require.NoError(t, err)

	// happens on the server, needs to be a request for this.
	signedCert, signedRaw, err := pki.SignUserCert(cp, keyPair.Cert, user)
	require.NoError(t, err)
	keyPair.SignedCert = signedCert
	keyPair.SignedCertPEM = signedRaw

	// create a test server and client to make sure the certs work.
	requireMutualTLSWorks(t, keyPair, cp)
}

func requireMutualTLSWorks(t *testing.T, clientKeypair *pki.KeyPair, cp pki.CertificateProvider) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "success!")
	}))

	serverTLSCerts, err := cp.TLSCertificates()
	require.NoError(t, err)

	caPool := x509.NewCertPool()

	for _, cert := range cp.ActiveCAs() {
		cert := cert
		caPool.AddCert(&cert)
	}

	server.TLS = &tls.Config{
		Certificates: serverTLSCerts,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caPool,
		MinVersion:   tls.VersionTLS12,
	}

	server.StartTLS()
	defer server.Close()

	// This will response with HTTP 200 OK and a body containing success!. We can now set up the client to trust the CA, and send a request to the server:

	clientTLSCert, err := clientKeypair.TLSCertificate()
	require.NoError(t, err)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{*clientTLSCert},
			ClientCAs:    caPool,
			RootCAs:      caPool,
			MinVersion:   tls.VersionTLS12,
		},
	}
	http := http.Client{
		Transport: transport,
	}

	resp, err := http.Get(server.URL)
	require.NoError(t, err)

	// If no errors occurred, we now have our success! response from the server, and can verify it:

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	body := strings.TrimSpace(string(respBodyBytes[:]))
	require.Equal(t, "success!", body)
}
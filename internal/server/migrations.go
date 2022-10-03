package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/infrahq/infra/api"
)

func (a *API) addRequestRewrites() {
	// all request migrations go here
}

func (a *API) addResponseRewrites() {
	// all response migrations go here
}

func (a *API) addRewrites() {
	a.addRequestRewrites()
	a.addResponseRewrites()
}

// addRedirects for API endpoints that have moved to a different path
func (a *API) addRedirects() {
}

func (a *API) deprecatedRoutes(noAuthnNoOrg *routeGroup) {
	// CLI clients before v0.14.4 rely on sign-up being false to continue with login
	type SignupEnabledResponse struct {
		Enabled bool `json:"enabled"`
	}

	add(a, noAuthnNoOrg, http.MethodGet, "/api/signup", route[api.EmptyRequest, *SignupEnabledResponse]{
		handler: func(c *gin.Context, _ *api.EmptyRequest) (*SignupEnabledResponse, error) {
			return &SignupEnabledResponse{Enabled: false}, nil
		},
		routeSettings: routeSettings{
			omitFromTelemetry: true,
			omitFromDocs:      true,
			txnOptions:        &sql.TxOptions{ReadOnly: true},
		},
	})
}

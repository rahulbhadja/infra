package access

import (
	"net/http"

	"github.com/infrahq/infra/internal/server/data"
	"github.com/infrahq/infra/internal/server/models"
)

const RequestContextKey = "requestContext"

// RequestContext stores the http.Request, and values derived from the request
// like the authenticated user. It also provides a database transaction.
type RequestContext struct {
	Request       *http.Request
	DBTxn         *data.Transaction
	Authenticated Authenticated

	// DataDB is the full database connection pool that can be used to
	// start transactions. Most routes should use DBTxn and should not use
	// DataDB directly.
	DataDB *data.DB
}

// Authenticated stores data about the authenticated user. If the AccessKey or
// User are nil, it indicates that no user was authenticated.
type Authenticated struct {
	AccessKey    *models.AccessKey
	User         *models.Identity
	Organization *models.Organization
}

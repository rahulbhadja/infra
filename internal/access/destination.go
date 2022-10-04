package access

import (
	"github.com/gin-gonic/gin"

	"github.com/infrahq/infra/internal/server/data"
	"github.com/infrahq/infra/internal/server/models"
	"github.com/infrahq/infra/uid"
)

func CreateDestination(c *gin.Context, destination *models.Destination) error {
	roles := []string{models.InfraAdminRole, models.InfraConnectorRole}
	db, err := RequireInfraRole(c, roles...)
	if err != nil {
		return HandleAuthErr(err, "destination", "create", roles...)
	}

	return data.CreateDestination(db, destination)
}

func SaveDestination(rCtx RequestContext, destination *models.Destination) error {
	roles := []string{models.InfraAdminRole, models.InfraConnectorRole}
	if err := IsAuthorized(rCtx, roles...); err != nil {
		return HandleAuthErr(err, "destination", "update", roles...)
	}

	return data.UpdateDestination(rCtx.DBTxn, destination)
}

func GetDestination(c *gin.Context, id uid.ID) (*models.Destination, error) {
	db := getDB(c)
	return data.GetDestination(db, data.ByID(id))
}

func ListDestinations(c *gin.Context, uniqueID, name string, p *data.Pagination) ([]models.Destination, error) {
	db := getDB(c)
	return data.ListDestinations(db, p, data.ByOptionalUniqueID(uniqueID),
		data.ByOptionalName(name))
}

func DeleteDestination(c *gin.Context, id uid.ID) error {
	db, err := RequireInfraRole(c, models.InfraAdminRole)
	if err != nil {
		return HandleAuthErr(err, "destination", "delete", models.InfraAdminRole)
	}

	return data.DeleteDestination(db, id)
}

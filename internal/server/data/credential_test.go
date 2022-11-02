package data

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/infrahq/infra/internal"
	"github.com/infrahq/infra/internal/server/models"
)

func TestDeleteCredential(t *testing.T) {
	runDBTests(t, func(t *testing.T, db *DB) {
		cred := &models.Credential{
			IdentityID:   7145,
			PasswordHash: []byte("password-hash"),
		}

		err := CreateCredential(db, cred)
		assert.NilError(t, err)

		t.Run("success", func(t *testing.T) {
			tx := txnForTestCase(t, db, db.DefaultOrg.ID)

			err := DeleteCredential(tx, cred.ID)
			assert.NilError(t, err)

			_, err = GetCredential(tx, ByIdentityID(7145))
			assert.ErrorIs(t, err, internal.ErrNotFound)

			// Delete again to check idempotence
			err = DeleteCredential(tx, cred.ID)
			assert.NilError(t, err)
		})
		t.Run("delete not found", func(t *testing.T) {
			tx := txnForTestCase(t, db, db.DefaultOrg.ID)

			err := DeleteCredential(tx, 171717)
			assert.NilError(t, err)
		})
	})
}

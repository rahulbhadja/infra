package data

import (
	"io/ioutil"
	"os"
	"testing"

	"gorm.io/gorm"
	"gotest.tools/v3/assert"

	"github.com/infrahq/infra/internal/server/models"
	"github.com/infrahq/infra/internal/testing/patch"
)

// see loadSQL for setting up your own migration test
func Test202204111503(t *testing.T) {
	db := setupWithNoMigrations(t, func(db *gorm.DB) {
		loadSQL(t, db, "202204111503")
	})

	// TODO: call NewDB to simulate a real server starting, instead of calling migrate
	err := migrate(db)
	assert.NilError(t, err)

	ids, err := ListIdentities(db, ByName("steven@example.com"))
	assert.NilError(t, err)

	assert.Assert(t, len(ids) == 1)
}

func Test202204211705(t *testing.T) {
	db := setupWithNoMigrations(t, func(db *gorm.DB) {
		loadSQL(t, db, "202204211705")
	})

	// TODO: call NewDB to simulate a real server starting, instead of calling migrate
	err := migrate(db)
	assert.NilError(t, err)

	// check it still works
	settings, err := GetSettings(db)
	assert.NilError(t, err)

	assert.Assert(t, settings != nil)
	assert.Assert(t, settings.PrivateJWK[0] == '{') // unencrypted type is json string.

	// check the storage data
	type Settings struct {
		models.Model
		PrivateJWK []byte
	}
	rawSettings := Settings{}
	err = db.Model(rawSettings).Where("id = ?", settings.ID).First(&rawSettings).Error
	assert.NilError(t, err)

	assert.Assert(t, rawSettings.PrivateJWK[0] != '{')
}

func Test202206151027(t *testing.T) {
	db := setupWithNoMigrations(t, func(db *gorm.DB) {
		loadSQL(t, db, "202206151027")
	})

	provider := map[string]interface{}{
		"name":          "test",
		"url":           "example.com",
		"client_id":     "a",
		"client_secret": "b",
	}
	err := db.Table("providers").Create(provider).Error
	assert.NilError(t, err)

	// TODO: call NewDB to simulate a real server starting, instead of calling migrate
	err = migrate(db)
	assert.NilError(t, err)

	result := map[string]interface{}{}
	err = db.Table("providers").Take(&result).Error
	assert.NilError(t, err)
	assert.Equal(t, models.OktaKind.String(), result["kind"])
}

// loadSQL loads a sql file from disk by a file name matching the migration it's meant to test.
// To create a new file for testing a migration:
//
// 1. Start an infra server and perform operations with the CLI or API to
//    get the db in the state you want to test. You should capture the db state
//    before writing your migration. Make sure there are some relevant records
//    in the affected tables.
//
// 2. Connect up to the db to dump the data. If you're running sqlite in
//    kubernetes:
//
//   kubectl exec -it deployment/infra-server -- apk add sqlite
//   kubectl exec -it deployment/infra-server -- /usr/bin/sqlite3 /var/lib/infrahq/server/sqlite3.db
//
//   at the prompt, do:
//     .dump
//   Copy to output to a file with the same name of the migration and a .sql
//   extension in the migrationdata/ folder.
//
//   Or from a local db:
//
//   echo -e ".output dump.sql\n.dump" | sqlite3 sqlite3.db
//
// 3. Write the migration and test that it does what you expect. It can be helpful
//    to put any necessary guards in place to make sure the database is in the state
//    you expect. Sometimes failed migrations leave it in a broken state, and might
//    run when you don't expect, so defensive programming is helpful here.
//
// 4. Go back to the sql file:
//   - remove any SQL records that aren't relevant to the test
//   - blank out provider client ids and name
//   - remove any email addresses and replace with @example.com
//   - remove any provider_users redirect urls
//   - check any other sensitive fields are encrypted and the key isn't included in the database.
func loadSQL(t *testing.T, db *gorm.DB, filename string) {
	f, err := os.Open("migrationdata/" + filename + ".sql")
	assert.NilError(t, err)

	b, err := ioutil.ReadAll(f)
	assert.NilError(t, err)

	err = db.Exec(string(b)).Error
	assert.NilError(t, err)
}

func setupWithNoMigrations(t *testing.T, f func(db *gorm.DB)) *gorm.DB {
	driver, err := NewSQLiteDriver("file::memory:")
	assert.NilError(t, err)

	db, err := newRawDB(driver)
	assert.NilError(t, err)

	f(db)

	patch.ModelsSymmetricKey(t)
	setupLogging(t)

	return db
}

package test

import (
	"database/sql"
	"os"
	"path"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
)

func isPathExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}

	if os.IsExist(err) {
		return true
	}

	return false
}

func LoadDBTestFixtures(t *testing.T, db *sql.DB, dialect string, suiteName string, testName string) {
	if db == nil || len(suiteName) == 0 || len(testName) == 0 {
		return
	}

	dataPath := path.Join("testdata", suiteName, testName, "database")
	if !isPathExist(dataPath) {
		return
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect(dialect),
		testfixtures.Directory(dataPath),
	)
	assert.NoError(t, err)

	err = fixtures.Load()
	assert.NoError(t, err)
}

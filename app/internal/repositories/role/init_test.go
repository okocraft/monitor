package role_test

import (
	"log"
	"os"
	"testing"

	"github.com/okocraft/monitor/lib/testutils/testdb"
)

var testDB testdb.TestDB

func TestMain(m *testing.M) {
	db, err := testdb.NewTestDB(true)
	if err != nil {
		log.Fatal(err)
	}

	testDB = db
	code := m.Run()

	err = db.Cleanup()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

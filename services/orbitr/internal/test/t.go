//lint:file-ignore U1001 Ignore all unused code, thinks the code is unused because of the test skips
package test

import (
	"io"
	"testing"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/metriqorg/go/services/orbitr/internal/db2/schema"
	"github.com/metriqorg/go/services/orbitr/internal/ledger"
	"github.com/metriqorg/go/services/orbitr/internal/operationfeestats"
	tdb "github.com/metriqorg/go/services/orbitr/internal/test/db"
	"github.com/metriqorg/go/services/orbitr/internal/test/scenarios"
	"github.com/metriqorg/go/support/db"
	"github.com/metriqorg/go/support/render/hal"
)

// CoreSession returns a db.Session instance pointing at the gravity test database
func (t *T) CoreSession() *db.Session {
	return &db.Session{
		DB: t.CoreDB,
	}
}

// Finish finishes the test, logging any accumulated orbitr logs to the logs
// output
func (t *T) Finish() {
	logEntries := t.testLogs()
	operationfeestats.ResetState()

	for _, entry := range logEntries {
		logString, err := entry.String()
		if err != nil {
			t.T.Logf("Error from entry.String: %v", err)
		} else {
			t.T.Log(logString)
		}
	}
}

// OrbitRSession returns a db.Session instance pointing at the orbitr test
// database
func (t *T) OrbitRSession() *db.Session {
	return &db.Session{
		DB: t.OrbitRDB,
	}
}

func (t *T) loadScenario(scenarioName string, includeOrbitR bool) {
	gravityPath := scenarioName + "-core.sql"

	scenarios.Load(tdb.GravityURL(), gravityPath)

	if includeOrbitR {
		orbitrPath := scenarioName + "-orbitr.sql"
		scenarios.Load(tdb.OrbitRURL(), orbitrPath)
	}
}

// Scenario loads the named sql scenario into the database
func (t *T) Scenario(name string) ledger.Status {
	clearOrbitRDB(t.T, t.OrbitRDB)
	t.loadScenario(name, true)
	return t.LoadLedgerStatus()
}

// ScenarioWithoutOrbitR loads the named sql scenario into the database
func (t *T) ScenarioWithoutOrbitR(name string) ledger.Status {
	t.loadScenario(name, false)
	ResetOrbitRDB(t.T, t.OrbitRDB)
	return t.LoadLedgerStatus()
}

// ResetOrbitRDB sets up a new orbitr database with empty tables
func ResetOrbitRDB(t *testing.T, db *sqlx.DB) {
	clearOrbitRDB(t, db)
	_, err := schema.Migrate(db.DB, schema.MigrateUp, 0)
	if err != nil {
		t.Fatalf("could not run migrations up on test db: %v", err)
	}
}

func clearOrbitRDB(t *testing.T, db *sqlx.DB) {
	_, err := schema.Migrate(db.DB, schema.MigrateDown, 0)
	if err != nil {
		t.Fatalf("could not run migrations down on test db: %v", err)
	}
}

// UnmarshalPage populates dest with the records contained in the json-encoded page in r
func (t *T) UnmarshalPage(r io.Reader, dest interface{}) hal.Links {
	var env struct {
		Embedded struct {
			Records json.RawMessage `json:"records"`
		} `json:"_embedded"`
		Links struct {
			Self hal.Link `json:"self"`
			Next hal.Link `json:"next"`
			Prev hal.Link `json:"prev"`
		} `json:"_links"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")

	err = json.Unmarshal(env.Embedded.Records, dest)
	t.Require.NoError(err, "failed to decode records")

	return env.Links
}

// UnmarshalNext extracts and returns the next link
func (t *T) UnmarshalNext(r io.Reader) string {
	var env struct {
		Links struct {
			Next struct {
				Href string `json:"href"`
			} `json:"next"`
		} `json:"_links"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")
	return env.Links.Next.Href
}

// UnmarshalExtras extracts and returns extras content
func (t *T) UnmarshalExtras(r io.Reader) map[string]string {
	var resp struct {
		Extras map[string]string `json:"extras"`
	}

	err := json.NewDecoder(r).Decode(&resp)
	t.Require.NoError(err, "failed to decode page")

	return resp.Extras
}

// LoadLedgerStatus loads ledger state from the core db(or panicing on failure).
func (t *T) LoadLedgerStatus() ledger.Status {
	var next ledger.Status

	err := t.CoreSession().GetRaw(t.Ctx, &next, `
		SELECT
			COALESCE(MAX(ledgerseq), 0) as core_latest
		FROM ledgerheaders
	`)

	if err != nil {
		panic(err)
	}

	err = t.OrbitRSession().GetRaw(t.Ctx, &next, `
			SELECT
				COALESCE(MIN(sequence), 0) as history_elder,
				COALESCE(MAX(sequence), 0) as history_latest
			FROM history_ledgers
		`)

	if err != nil {
		panic(err)
	}

	return next
}

// retrieves entries from test logger instance
func (t *T) testLogs() []logrus.Entry {
	if t.EndLogTest == nil {
		return []logrus.Entry{}
	}

	return t.EndLogTest()
}

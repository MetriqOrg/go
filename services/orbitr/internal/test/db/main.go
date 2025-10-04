// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on orbitr and so should be able to be imported by
// any orbitr package.
package db

import (
	"fmt"
	"log"
	"testing"

	// pq enables postgres support
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	db "github.com/metriqorg/go/support/db/dbtest"
)

var (
	orbitrDB     *db.DB
	coreDB        *db.DB
	coreDBConn    *sqlx.DB
	orbitrDBConn *sqlx.DB
)

func orbitrPostgres(t *testing.T) *db.DB {
	if orbitrDB != nil {
		return orbitrDB
	}
	orbitrDB = db.Postgres(t)
	return orbitrDB
}

func corePostgres(t *testing.T) *db.DB {
	if coreDB != nil {
		return coreDB
	}
	coreDB = db.Postgres(t)
	return coreDB
}

func OrbitR(t *testing.T) *sqlx.DB {
	if orbitrDBConn != nil {
		return orbitrDBConn
	}

	orbitrDBConn = orbitrPostgres(t).Open()
	return orbitrDBConn
}

func OrbitRURL() string {
	if orbitrDB == nil {
		log.Panic(fmt.Errorf("OrbitR not initialized"))
	}
	return orbitrDB.DSN
}

func OrbitRROURL() string {
	if orbitrDB == nil {
		log.Panic(fmt.Errorf("OrbitR not initialized"))
	}
	return orbitrDB.RO_DSN
}

func Gravity(t *testing.T) *sqlx.DB {
	if coreDBConn != nil {
		return coreDBConn
	}
	coreDBConn = corePostgres(t).Open()
	return coreDBConn
}

func GravityURL() string {
	if coreDB == nil {
		log.Panic(fmt.Errorf("Gravity not initialized"))
	}
	return coreDB.DSN
}

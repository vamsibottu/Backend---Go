package pg

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var (
	pgxorm     *xorm.Engine
	oncepgxorm sync.Once
)

var (
	pgdb     *sql.DB
	oncepgdb sync.Once
)

func InitXorm() *xorm.Engine {

	oncepgxorm.Do(func() {
		psqlInfo := fmt.Sprintf("postgresql://localhost:5432/postgres?sslmode=disable")

		var err error
		pgxorm, err = xorm.NewEngine("postgres", psqlInfo)
		if err != nil {
			log.Error("error connecting to pg xorm:", err)
			return
		}

		//TODO: need to make these configurable
		pgxorm.ShowSQL(true)
		pgxorm.SetMaxOpenConns(5)
		pgxorm.SetMaxIdleConns(5)
	})

	return pgxorm
}

func Initdb() *sql.DB {

	oncepgdb.Do(func() {
		psqlInfo := fmt.Sprintf("postgresql://localhost:5432/postgres?sslmode=disable")

		var err error
		pgdb, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Error("error connecting to pg db:", err)
			return
		}

		//TODO: need to make these configurable
		pgdb.SetMaxOpenConns(5)
		pgdb.SetMaxIdleConns(5)

	})
	return pgdb
}

package fields

import (
	"database/sql"

	"github.com/Datamigration/internal/platform/db/pg"
	"github.com/go-xorm/xorm"
)

var pqdb *sql.DB
var pqxorm *xorm.Engine

func init() {
	pqxorm = pg.InitXorm()

}

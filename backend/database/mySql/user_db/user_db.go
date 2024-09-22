package userdb

import "database/sql"

var (
	Clinte *sql.DB
)

func init() {

	dataSourceName :=
		sql.Open("mySQL", dataSourceName)
}

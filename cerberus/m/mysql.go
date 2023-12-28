package m

import (
	"database/sql"

	"github.com/iegad/gox/cerberus/conf"
	"github.com/iegad/gox/frm/db"
	"github.com/iegad/gox/frm/log"
)

var Mysql *sql.DB

func InitMysql() {
	var err error
	Mysql, err = db.NewSql(conf.Instance.Mysql)
	if err != nil {
		log.Fatal(err)
	}
}

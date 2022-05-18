package psql

import (
	"github.com/go-pg/pg/v10"
	"go-project/pkg/setting"
)

var (
	Pgdb *pg.DB
)

// 初始化psql
func Setup() {
	opt, err := pg.ParseURL(setting.PostgresqlDbSetting.Url)
	if err != nil {
		panic(err)
	}
	Pgdb = pg.Connect(opt)
}

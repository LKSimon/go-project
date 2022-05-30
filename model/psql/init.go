package psql

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
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
	Pgdb.AddQueryHook(&debugger{})
}

type debugger struct{}

func (d *debugger) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
	q, err := evt.FormattedQuery()
	if err != nil {
		return ctx, err
	}

	if _, ok := ctx.(*gin.Context); ok {
		log.Printf("pg query:%s", string(q))
	}

	return ctx, nil
}

func (d *debugger) AfterQuery(context.Context, *pg.QueryEvent) error {
	return nil
}

package main

import (
	"fmt"

	"go-project/data"
	"go-project/model/psql"
	"go-project/pkg/setting"
)

func init() {
	setting.Setup("dev")
	psql.Setup()
}

func main() {
	gameDepotModel := data.NewGameDepot(psql.Pgdb, setting.PostgresqlDbSetting.Schema)
	gameDepot, err := gameDepotModel.Get("120259084289")
	if err != nil {
		panic(err)
	}
	fmt.Printf("gameDepot: %+v\n", gameDepot)
}

package data

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"go-project/model/psql"
)

type GameDepot struct {
	*pg.DB
	Schema string
}

func NewGameDepot(db *pg.DB, schema string) *GameDepot {
	return &GameDepot{
		DB:     db,
		Schema: schema,
	}
}

func (g *GameDepot) Get(id string) (*psql.GameDepot, error) {
	model := psql.GameDepot{}
	tb := fmt.Sprintf("%s.tb_game_depot", g.Schema)
	sql := fmt.Sprintf(`SELECT * FROM %s WHERE depot_id = ?`, tb)
	_, err := g.DB.Model(&model).Table(tb).QueryOne(&model, sql, id)
	return &model, err
}

func (g *GameDepot) List(offset, limit int) ([]psql.GameDepot, error) {
	model := []psql.GameDepot{}
	tb := fmt.Sprintf("%s.tb_game_depot", g.Schema)
	err := g.DB.Model(&model).Table(tb).Offset(offset).Limit(limit).Select()
	return model, err
}

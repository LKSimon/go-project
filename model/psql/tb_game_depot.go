package psql

// 表名：tb_game_depot
type GameDepot struct {
	tableName           struct{} `pg:"_"`        // 使用动态表名，通过Table()函数指定
	DepotId             string   `pg:"depot_id"` // 目录id (主键)
	DepotName           string   `pg:"depot_name"`
	PublicToken         string   `pg:"public_token"`
	GameCode            string   `pg:"game_code"`
	GameName            string   `pg:"game_name"`
	UpdateDate          string   `pg:"update_date"`
	ArthubCode          string   `pg:"arthub_code"`
	Type                int      `pg:"type"` // 默认 1表示arthub，2表示google drive
	GoogleServceAccount string   `pg:"google_servce_account"`
}

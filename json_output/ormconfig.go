package json_output

import (
	"encoding/json"
	"io/ioutil"
)

/*
 * TypeOrm configuration definition
 */
type OrmCongfig struct {
	Type        string   `json:"type"`
	Host        string   `json:"host"`
	Port        int16    `json:"port"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Database    string   `json:"database"`
	Synchronize bool     `json:"synchronize"`
	Logging     bool     `json:"logging"`
	Entity      []string `json:"entity"`
	Migrations  []string `json:"migrations"`
	Subscribers []string `json:"subscribers"`
}

func (o OrmCongfig) Write() {
	file, _ := json.MarshalIndent(o, "", "  ")
	_ = ioutil.WriteFile("ormconfig.json", file, permissions)
}

func NewOrmConfig() OrmCongfig {
	config := OrmCongfig{}
	config.Type = "postgress"
	config.Host = "localhost"
	config.Port = 5432
	config.Username = "YourUser"
	config.Password = "YourPassword"
	config.Database = "Database"
	config.Synchronize = false
	config.Logging = false
	config.Entity = []string{"src/entity/**/*.ts"}
	config.Migrations = []string{"src/database/migrations/**/*.ts"}
	config.Subscribers = []string{"src/subscriber/**/*.ts"}

	return config
}

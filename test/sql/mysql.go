package gosql

import "database/sql"

type MySQLConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	DBName   string `json:"dbname"`
	Params   string `json:"params"`
}

func NewDBMySQL(conf MySQLConfig) *sql.DB {
	// set Data Source Name
	var dsn string

	// username:password@protocol(address)/dbname?param=value
	dsn = conf.UserName + ":" + conf.Password + "@" + conf.Protocol +
		"(" + conf.Address + ")/" + conf.DBName
	if conf.Params != "" {
		dsn += "?" + conf.Params
	}

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic("failed to open database")
	}

	return db
}

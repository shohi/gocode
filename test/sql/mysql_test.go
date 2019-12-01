package gosql

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// NOTE: Make sure MySQL server is running
// 1. docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=passwd -d mysql:8.0.18
// FIXME: use correct default config
func newTestMySQLConfig() MySQLConfig {
	return MySQLConfig{
		UserName: "root",
		Password: "passwd",
		Protocol: "tcp",
		Address:  "localhost",
		DBName:   "",
	}
}

func newTestDB() *sql.DB {
	conf := newTestMySQLConfig()
	return NewDBMySQL(conf)
}

func TestDrivers(t *testing.T) {
	log.Println(sql.Drivers())
}

func TestPing(t *testing.T) {
	db := newTestDB()
	log.Println(db.Ping())
}

func TestQuery(t *testing.T) {
	db := newTestDB()
	rows, err := db.Query("show databases")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		log.Printf("name is %s\n", name)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
}

func TestDatabaseStatus(t *testing.T) {
	db := newTestDB()
	log.Println(db.Stats())
}

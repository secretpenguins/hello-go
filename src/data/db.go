package data

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/coopernurse/gorp"
	"fmt"
	"log"
	"github.com/bradfitz/gomemcache/memcache"
	"config"
)

var dbMap *gorp.DbMap
var globalDb *sql.DB
var mc *memcache.Client

func init() {
	Setup()
	//defer dbMap.Db.Close()
}

func Setup() {
	log.Println("DB Init");
	//go2:password@tcp(localhost:3306)/blog
	config := config.GetConfig()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbDatabase);
	log.Println(connectionString)

	db, _ := sql.Open("mysql", connectionString)

	globalDb = db
	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbMap.AddTableWithName(Post{}, "Posts").SetKeys(true, "postId")
	dbMap.AddTableWithName(Login{}, "Logins").SetKeys(true, "LoginId")

	mc = memcache.New(config.MemcachePath)
}

func GetDbMap() *gorp.DbMap {
	return dbMap;
}
package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
)

var DbEngin *xorm.Engine

func SetDB()  {
	db, err := xorm.NewEngine(viper.GetString("db.drive_name"), viper.GetString("db.link"))
	if err != nil {
		log.Fatal(err.Error())
	}
	db.ShowSQL(viper.GetBool("db.is_show_sql"))
	db.SetMaxOpenConns(viper.GetInt("db.max_connection"))
}

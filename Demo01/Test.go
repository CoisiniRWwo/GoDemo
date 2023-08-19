package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func init() {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local ", user, password, ip, port, database)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印 sql
		//SkipDefaultTransaction: true, //禁用事务
	})
	if err != nil {
		fmt.Println(err)
	}
}

package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"	
)

var DB *gorm.DB

type DBConfig struct {
	ServerName string
	Port int
	User string
	Password string
	DB string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig {
		ServerName: "localhost",
		Port: 3306,
		User: "root",
		Password: "",
		DB: "apidb",
	}
	
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", 
		dbConfig.User, 
		dbConfig.Password, 
		dbConfig.ServerName, 
		dbConfig.Port,
		dbConfig.DB,
	)
}

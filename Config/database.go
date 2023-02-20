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

func BuildDBConfig(env string) *DBConfig {
	var dbConfig DBConfig;
	if (env == "prod") {
		dbConfig = DBConfig {
			ServerName: "localhost",
			Port: 3306,
			User: "root",
			Password: "",
			DB: "apidb",
		}
	} else if (env == "test") {
		dbConfig = DBConfig {
			ServerName: "localhost",
			Port: 3306,
			User: "root",
			Password: "",
			DB: "apidb_test",
		}
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

package main

import (
	"api/Config"
	"api/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	fmt.Println("Connecting to DB...")

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig())) // To setup MySQL DB.

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close() // Keep

	fmt.Println("Successfully connected to database.")

	r := Routes.Setup()
	r.Run()
}


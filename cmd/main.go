package main

import (
	"FirstWeek/Config"
	"FirstWeek/Internal/Repository"
	"FirstWeek/Internal/Services"
	"FirstWeek/Internal/adapter/api"
	"FirstWeek/Internal/adapter/db"
	"fmt"
)

func main() {

	var configuration Config.Configurations
	config, err := Config.SetUpViper(configuration)
	if err != nil {
		fmt.Println("Failed to read Config file")
	}
	con, err := db.NewDBConnection(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	tranRepo := Repository.NewDefaultRepository(con)
	tranSer := Services.NewDefaultService(tranRepo)
	api.NewTransactionController(tranSer, config)
}

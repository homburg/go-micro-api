package main

import (
	"fmt"

	"github.com/byblix/go-micro/config"
	"github.com/byblix/go-micro/routes"
	"github.com/spf13/viper"
)

func main() {
	// Set env
	config.SetEnvironment()
	fmt.Println(viper.Get("dev.ENV"))
	// Init routes from init_routes.go
	routes.InitRoutes()

	// scripts
	// scripts.initScripts()
}

func initScripts() {
	// scripts.WithdrawalsToCSV("withdrawals")
	// scripts.ProfilesToCSV("profiles")
}

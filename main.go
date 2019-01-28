package main

import (
	"github.com/byblix/go-micro/config"
	"github.com/byblix/go-micro/routes"
)

func main() {
	// Set env
	config.SetEnvironment("hello")
	// Init routes from init_routes.go
	routes.InitRoutes()

	// scripts
	// scripts.initScripts()
}

func initScripts() {
	// scripts.WithdrawalsToCSV("withdrawals")
	// scripts.ProfilesToCSV("profiles")
}

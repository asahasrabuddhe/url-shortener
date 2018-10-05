package main

import (
	"flag"
	"github.com/asahasrabuddhe/go-api-base/server"
	"github.com/asahasrabuddhe/url-shortener/url/datastore"
	"github.com/asahasrabuddhe/url-shortener/url/helpers"
	"github.com/asahasrabuddhe/url-shortener/url/routes"
)

func main() {
	configPath := flag.String("config-path", "", "Path to the config file location (Required)")
	configFileName := flag.String("config-name", "config", "Name of the config file without the extension (Required)")
	flag.Parse()

	server.Init(*configPath, *configFileName)

	datastore.DB = datastore.NewDatabase("ajitem.db")

	routes.CreateUrlRoutes()

	err := server.Start()
	helpers.HandleError(err)
}

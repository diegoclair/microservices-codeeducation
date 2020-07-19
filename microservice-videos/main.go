package main

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/data"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/service"
)

func main() {
	logger.Info("Reading the intial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}
	svc := service.New(db)
	server := server.InitServer(svc)
	logger.Info("About to start the application...")

	if err := server.Run(":3000"); err != nil {
		panic(err)
	}
}

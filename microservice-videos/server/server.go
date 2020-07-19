package server

import (
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	_ = service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &controller{})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	return srv
}

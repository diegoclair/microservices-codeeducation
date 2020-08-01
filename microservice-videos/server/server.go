package server

import (
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server/routes/categoryroute"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	categoryController *categoryroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		categoryController: categoryroute.NewController(svm.CategoryService(svc)),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	return srv
}

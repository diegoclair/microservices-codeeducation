package server

import (
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server/routes/categoryroute"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/service"
	"github.com/gin-gonic/gin"
)

type routes struct {
	categoryController *categoryroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	srv := gin.Default()

	return setupRoutes(srv, &routes{
		categoryController: categoryroute.NewController(svm.CategoryService(svc)),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, r *routes) *gin.Engine {

	categoryroute.NewRouter(r.categoryController, srv).RegisterRoutes()

	return srv
}

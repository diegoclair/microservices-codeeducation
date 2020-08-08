package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/data/factory"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server/routes/categoryroute"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/service"
	"github.com/gin-gonic/gin"
)

type routes struct {
	categoryController *categoryroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	mapper := mapper.New()
	svm := service.NewServiceManager()
	srv := gin.Default()

	categoryService := svm.CategoryService(svc)

	fakedata(categoryService)

	return setupRoutes(srv, &routes{
		categoryController: categoryroute.NewController(categoryService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, r *routes) *gin.Engine {

	categoryroute.NewRouter(r.categoryController, srv).RegisterRoutes()

	return srv
}

func fakedata(c contract.CategoryService) {

	logger.Info("Create fake data")
	err := factory.FakeData(c)
	if err != nil {
		panic(err)
	}
}

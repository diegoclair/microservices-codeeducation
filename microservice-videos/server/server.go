package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/data/factory"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server/routes/categoryroute"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server/routes/genreroute"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/service"
	"github.com/gin-gonic/gin"
)

type routes struct {
	categoryController *categoryroute.Controller
	genreController    *genreroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	mapper := mapper.New()
	svm := service.NewServiceManager()
	srv := gin.Default()

	categoryService := svm.CategoryService(svc)
	genreService := svm.GenreService(svc)

	fakedata(categoryService, genreService)

	return setupRoutes(srv, &routes{
		categoryController: categoryroute.NewController(categoryService, mapper),
		genreController:    genreroute.NewController(genreService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, r *routes) *gin.Engine {

	categoryroute.NewRouter(r.categoryController, srv).RegisterRoutes()
	genreroute.NewRouter(r.genreController, srv).RegisterRoutes()

	return srv
}

func fakedata(c contract.CategoryService, g contract.GenreService) {

	logger.Info("Create fake data")
	err := factory.FakeData(c, g)
	if err != nil {
		panic(err)
	}
}

package rest

import (
	"fmt"
	"os"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/application/factory"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/application/rest/routes/categoryroute"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/application/rest/routes/genreroute"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	dataFactory "github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/data/factory"

	"github.com/gin-gonic/gin"
)

type routes struct {
	categoryController *categoryroute.Controller
	genreController    *genreroute.Controller
}

//StartRestServer starts the restServer
func StartRestServer() {
	server := initServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	logger.Info("About to start the application...")

	if err := server.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}

}

//initServer to initialize the server
func initServer() *gin.Engine {

	factory := factory.GetDomainServices()
	srv := gin.Default()

	fakedata(factory.CategoryService, factory.GenreService)

	return registerRoutes(srv, &routes{
		categoryController: categoryroute.NewController(factory.CategoryService, factory.Mapper),
		genreController:    genreroute.NewController(factory.GenreService, factory.Mapper),
	})
}

//registerRoutes - Register and instantiate the routes
func registerRoutes(srv *gin.Engine, r *routes) *gin.Engine {

	categoryroute.NewRouter(r.categoryController, srv).RegisterRoutes()
	genreroute.NewRouter(r.genreController, srv).RegisterRoutes()

	return srv
}

func fakedata(c contract.CategoryService, g contract.GenreService) {

	logger.Info("Create fake data")
	err := dataFactory.FakeData(c, g)
	if err != nil {
		panic(err)
	}
}

package factory

import (
	"log"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/domain/service"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/config"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/data"
)

//Services is the factory to all serrvices
type Services struct {
	Cfg             *config.EnvironmentVariables
	Mapper          mapper.Mapper
	CategoryService contract.CategoryService
	GenreService    contract.GenreService
}

var (
	instance *Services
	once     sync.Once
)

//GetDomainServices to get instace of all services
func GetDomainServices() *Services {

	once.Do(func() {

		data, err := data.Connect()
		if err != nil {
			log.Fatalf("Error to connect data repositories: %v", err)
		}

		cfg := config.GetConfigEnvironment()
		svc := service.New(data, cfg)
		svm := service.NewServiceManager()
		mapper := mapper.New()

		instance = &Services{
			Cfg:             cfg,
			Mapper:          mapper,
			CategoryService: svm.CategoryService(svc),
			GenreService:    svm.GenreService(svc),
		}

	})

	return instance
}

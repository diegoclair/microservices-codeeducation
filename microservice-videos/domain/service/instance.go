package service

import (
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/infra/config"
)

// Service holds the domain service repositories
type Service struct {
	dm  contract.DataManager
	cfg *config.EnvironmentVariables
}

// New returns a new domain Service instance
func New(dm contract.DataManager, cfg *config.EnvironmentVariables) *Service {
	svc := new(Service)
	svc.dm = dm
	svc.cfg = cfg

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	CategoryService(svc *Service) contract.CategoryService
	GenreService(svc *Service) contract.GenreService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) CategoryService(svc *Service) contract.CategoryService {
	return newCategoryService(svc)
}

func (s *serviceManager) GenreService(svc *Service) contract.GenreService {
	return newGenreService(svc)
}

package service

import "github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"

// Service holds the domain service repositories
type Service struct {
	db contract.RepoManager
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

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

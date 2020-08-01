package service

import "github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"

type categoryService struct {
	svc *Service
}

//newCategoryService return a new instance of the service
func newCategoryService(svc *Service) contract.CategoryService {
	return &categoryService{
		svc: svc,
	}
}

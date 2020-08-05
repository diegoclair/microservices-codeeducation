package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
)

type categoryService struct {
	svc *Service
}

//newCategoryService return a new instance of the service
func newCategoryService(svc *Service) contract.CategoryService {
	return &categoryService{
		svc: svc,
	}
}

func (s *categoryService) GetCategories() ([]*entity.Category, resterrors.RestErr) {

	return s.svc.db.Category().GetCategories()
}

func (s *categoryService) GetCategoryByID(id int64) (*entity.Category, resterrors.RestErr) {

	return s.svc.db.Category().GetCategoryByID(id)
}

func (s *categoryService) CreateCategory(category entity.Category) (*entity.Category, resterrors.RestErr) {

	err := category.Validate()
	if err != nil {
		return nil, err
	}

	return s.svc.db.Category().CreateCategory(category)
}

func (s *categoryService) UpdateCategoryByID(id int64, category entity.Category) resterrors.RestErr {

	err := category.Validate()
	if err != nil {
		return err
	}

	return s.svc.db.Category().UpdateCategoryByID(id, category)
}

func (s *categoryService) DeleteCategoryByID(id int64) resterrors.RestErr {

	return s.svc.db.Category().DeleteCategoryByID(id)
}
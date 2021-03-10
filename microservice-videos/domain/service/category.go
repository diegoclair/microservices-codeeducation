package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/go_utils-lib/validstruct"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/domain/entity"
	uuid "github.com/satori/go.uuid"
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

func (s *categoryService) GetCategories() (*[]entity.Category, resterrors.RestErr) {

	return s.svc.dm.MySQL().Category().GetCategories()
}

func (s *categoryService) GetCategoryByUUID(uuid string) (*entity.Category, resterrors.RestErr) {

	return s.svc.dm.MySQL().Category().GetCategoryByUUID(uuid)
}

func (s *categoryService) CreateCategory(category entity.Category) (*entity.Category, resterrors.RestErr) {

	category.UUID = uuid.NewV4().String()

	err := validstruct.ValidateStruct(category)
	if err != nil {
		return nil, err
	}

	err = s.svc.dm.MySQL().Category().CreateCategory(category)
	if err != nil {
		return nil, err
	}

	categoryData, err := s.svc.dm.MySQL().Category().GetCategoryByUUID(category.UUID)
	if err != nil {
		return nil, err
	}

	return categoryData, nil
}

func (s *categoryService) UpdateCategoryByID(uuid string, category entity.Category) resterrors.RestErr {

	err := validstruct.ValidateStruct(category)
	if err != nil {
		return err
	}

	return s.svc.dm.MySQL().Category().UpdateCategoryByID(uuid, category)
}

func (s *categoryService) DeleteCategoryByID(uuid string) resterrors.RestErr {

	return s.svc.dm.MySQL().Category().DeleteCategoryByID(uuid)
}

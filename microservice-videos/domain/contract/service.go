package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
)

// CategoryService holds a user service operations
type CategoryService interface {
	GetCategories() (*[]entity.Category, resterrors.RestErr)
	GetCategoryByUUID(uuid string) (*entity.Category, resterrors.RestErr)
	CreateCategory(category entity.Category) (*entity.Category, resterrors.RestErr)
	UpdateCategoryByID(uuid string, category entity.Category) resterrors.RestErr
	DeleteCategoryByID(uuid string) resterrors.RestErr
}

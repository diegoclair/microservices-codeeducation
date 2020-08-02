package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
)

// CategoryService holds a user service operations
type CategoryService interface {
	GetCategories() ([]entity.Category, resterrors.RestErr)
	GetCategoryByID(id int64) (entity.Category, resterrors.RestErr)
	UpdateCategoryByID(id int64) resterrors.RestErr
	CreateCategory(category entity.Category) resterrors.RestErr
}

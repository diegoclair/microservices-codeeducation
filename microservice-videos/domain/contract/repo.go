package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Category() CategoryRepo
}

// CategoryRepo data set
type CategoryRepo interface {
	GetCategories() (*[]entity.Category, resterrors.RestErr)
	GetCategoryByUUID(uuid string) (*entity.Category, resterrors.RestErr)
	CreateCategory(category entity.Category) resterrors.RestErr
	UpdateCategoryByID(uuid string, category entity.Category) resterrors.RestErr
	DeleteCategoryByID(uuid string) resterrors.RestErr
}

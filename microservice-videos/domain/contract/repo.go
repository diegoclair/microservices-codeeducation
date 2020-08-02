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
	GetCategories() ([]entity.Category, resterrors.RestErr)
	GetCategoryByID(id int64) (entity.Category, resterrors.RestErr)
}

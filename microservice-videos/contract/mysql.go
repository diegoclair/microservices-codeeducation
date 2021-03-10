package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/domain/entity"
)

//MySQLRepo defines the repository aggregator interface
type MySQLRepo interface {
	Category() CategoryRepo
	Genre() GenreRepo
}

// CategoryRepo data set
type CategoryRepo interface {
	GetCategories() (*[]entity.Category, resterrors.RestErr)
	GetCategoryByUUID(uuid string) (*entity.Category, resterrors.RestErr)
	CreateCategory(category entity.Category) resterrors.RestErr
	UpdateCategoryByID(uuid string, category entity.Category) resterrors.RestErr
	DeleteCategoryByID(uuid string) resterrors.RestErr
}

// GenreRepo data set
type GenreRepo interface {
	GetGenres() (*[]entity.Genre, resterrors.RestErr)
	GetGenreByUUID(uuid string) (*entity.Genre, resterrors.RestErr)
	CreateGenre(category entity.Genre) resterrors.RestErr
	UpdateGenreByID(uuid string, category entity.Genre) resterrors.RestErr
	DeleteGenreByID(uuid string) resterrors.RestErr
}

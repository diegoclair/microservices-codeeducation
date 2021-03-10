package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/domain/entity"
)

// CategoryService holds a user service operations
type CategoryService interface {
	GetCategories() (*[]entity.Category, resterrors.RestErr)
	GetCategoryByUUID(uuid string) (*entity.Category, resterrors.RestErr)
	CreateCategory(category entity.Category) (*entity.Category, resterrors.RestErr)
	UpdateCategoryByID(uuid string, category entity.Category) resterrors.RestErr
	DeleteCategoryByID(uuid string) resterrors.RestErr
}

// GenreService holds a user service operations
type GenreService interface {
	GetGenres() (*[]entity.Genre, resterrors.RestErr)
	GetGenreByUUID(uuid string) (*entity.Genre, resterrors.RestErr)
	CreateGenre(category entity.Genre) (*entity.Genre, resterrors.RestErr)
	UpdateGenreByID(uuid string, category entity.Genre) resterrors.RestErr
	DeleteGenreByID(uuid string) resterrors.RestErr
}

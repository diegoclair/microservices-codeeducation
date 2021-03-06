package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/go_utils-lib/validstruct"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type genreService struct {
	svc *Service
}

//newGenreService return a new instance of the service
func newGenreService(svc *Service) contract.GenreService {
	return &genreService{
		svc: svc,
	}
}

func (s *genreService) GetGenres() (*[]entity.Genre, resterrors.RestErr) {

	return s.svc.dm.MySQL().Genre().GetGenres()
}

func (s *genreService) GetGenreByUUID(uuid string) (*entity.Genre, resterrors.RestErr) {

	return s.svc.dm.MySQL().Genre().GetGenreByUUID(uuid)
}

func (s *genreService) CreateGenre(genre entity.Genre) (*entity.Genre, resterrors.RestErr) {

	genre.UUID = uuid.NewV4().String()

	err := validstruct.ValidateStruct(genre)
	if err != nil {
		return nil, err
	}

	err = s.svc.dm.MySQL().Genre().CreateGenre(genre)
	if err != nil {
		return nil, err
	}

	genreData, err := s.svc.dm.MySQL().Genre().GetGenreByUUID(genre.UUID)
	if err != nil {
		return nil, err
	}

	return genreData, nil
}

func (s *genreService) UpdateGenreByID(uuid string, genre entity.Genre) resterrors.RestErr {

	err := validstruct.ValidateStruct(genre)
	if err != nil {
		return err
	}

	return s.svc.dm.MySQL().Genre().UpdateGenreByID(uuid, genre)
}

func (s *genreService) DeleteGenreByID(uuid string) resterrors.RestErr {

	return s.svc.dm.MySQL().Genre().DeleteGenreByID(uuid)
}

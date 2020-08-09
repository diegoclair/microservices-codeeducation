package factory

import (
	"fmt"
	"strconv"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
)

//FakeData to include fake data on database
func FakeData(c contract.CategoryService, g contract.GenreService) resterrors.RestErr {

	logger.Info("Creating category fake data...")
	err := categoryData(c)
	if err != nil {
		return err
	}

	logger.Info("Creating genre fake data...")
	err = genreData(g)
	if err != nil {
		return err
	}

	return nil
}

func categoryData(s contract.CategoryService) resterrors.RestErr {

	_, err := s.GetCategories()
	if err != nil && err.StatusCode() != 404 {
		logger.Error("FakeGetCategories: ", fmt.Errorf(err.Message()))
		return err
	}
	if err != nil && err.StatusCode() == 404 {
		//if the err code is 404, it means that we have no categories in database, then we'll create
		for i := 0; i < 100; i++ {
			fake := entity.Category{}
			fake.Name = "Test name " + strconv.Itoa(i+1)
			fake.Description = "Description " + strconv.Itoa(i+1)
			_, err = s.CreateCategory(fake)
			if err != nil {
				logger.Error("FakeCreateCategory: ", fmt.Errorf(err.Message()))
				return err
			}
		}
	}

	return nil
}

func genreData(s contract.GenreService) resterrors.RestErr {

	_, err := s.GetGenres()
	if err != nil && err.StatusCode() != 404 {
		logger.Error("FakeGetGenres: ", fmt.Errorf(err.Message()))
		return err
	}
	if err != nil && err.StatusCode() == 404 {
		//if the err code is 404, it means that we have no categories in database, then we'll create
		for i := 0; i < 100; i++ {
			fake := entity.Genre{}
			fake.Name = "Test name " + strconv.Itoa(i+1)
			_, err = s.CreateGenre(fake)
			if err != nil {
				logger.Error("FakeCreateGenre: ", fmt.Errorf(err.Message()))
				return err
			}
		}
	}

	return nil
}

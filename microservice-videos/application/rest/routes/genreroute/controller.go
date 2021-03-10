package genreroute

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/application/rest/viewmodel"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/contract"
	"github.com/diegoclair/microservices-codeeducation/microservice-videos/domain/entity"
	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds genre handlers
type Controller struct {
	genreService contract.GenreService
	mapper       mapper.Mapper
}

//NewController to handle requests
func NewController(genreService contract.GenreService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			genreService: genreService,
			mapper:       mapper,
		}
	})
	return instance
}

func (c *Controller) handleGetGenres(ctx echo.Context) error {

	genres, err := c.genreService.GetGenres()
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := []viewmodel.Genre{}
	mapErr := c.mapper.From(*genres).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleGetGenreByID(ctx echo.Context) error {

	uuid := ctx.Param("genre_id")

	genres, err := c.genreService.GetGenreByUUID(uuid)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := viewmodel.Genre{}
	mapErr := c.mapper.From(*genres).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleCreateGenre(ctx echo.Context) error {

	var input entity.Genre

	jsonErr := ctx.Bind(&input)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		return ctx.JSON(err.StatusCode(), err)
	}

	genre, err := c.genreService.CreateGenre(input)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := viewmodel.Genre{}
	mapErr := c.mapper.From(*genre).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) handleUpdateGenreByID(ctx echo.Context) error {

	uuid := ctx.Param("genre_id")

	var genre entity.Genre
	jsonErr := ctx.Bind(&genre)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		return ctx.JSON(err.StatusCode(), err)
	}

	err := c.genreService.UpdateGenreByID(uuid, genre)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *Controller) handleDeleteGenreByID(ctx echo.Context) error {

	uuid := ctx.Param("genre_id")

	err := c.genreService.DeleteGenreByID(uuid)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

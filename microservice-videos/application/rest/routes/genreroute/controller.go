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

	"github.com/gin-gonic/gin"
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

func (c *Controller) handleGetGenres(ctx *gin.Context) {

	genres, err := c.genreService.GetGenres()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	response := []viewmodel.Genre{}
	mapErr := c.mapper.From(*genres).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		ctx.JSON(err.StatusCode(), err)
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleGetGenreByID(ctx *gin.Context) {

	uuid := ctx.Param("genre_id")

	genres, err := c.genreService.GetGenreByUUID(uuid)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	response := viewmodel.Genre{}
	mapErr := c.mapper.From(*genres).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		ctx.JSON(err.StatusCode(), err)
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleCreateGenre(ctx *gin.Context) {

	var input entity.Genre

	jsonErr := ctx.ShouldBindJSON(&input)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		ctx.JSON(err.StatusCode(), err)
		return
	}

	genre, err := c.genreService.CreateGenre(input)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	response := viewmodel.Genre{}
	mapErr := c.mapper.From(*genre).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		ctx.JSON(err.StatusCode(), err)
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) handleUpdateGenreByID(ctx *gin.Context) {

	uuid := ctx.Param("genre_id")

	var genre entity.Genre
	jsonErr := ctx.ShouldBindJSON(&genre)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		ctx.JSON(err.StatusCode(), err)
		return
	}

	err := c.genreService.UpdateGenreByID(uuid, genre)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (c *Controller) handleDeleteGenreByID(ctx *gin.Context) {

	uuid := ctx.Param("genre_id")

	err := c.genreService.DeleteGenreByID(uuid)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

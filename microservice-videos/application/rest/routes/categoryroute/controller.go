package categoryroute

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

//Controller holds category handlers
type Controller struct {
	categoryService contract.CategoryService
	mapper          mapper.Mapper
}

//NewController to handle requests
func NewController(categoryService contract.CategoryService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			categoryService: categoryService,
			mapper:          mapper,
		}
	})
	return instance
}

func (c *Controller) handleGetCategories(ctx echo.Context) error {

	categories, err := c.categoryService.GetCategories()
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := []viewmodel.Category{}
	mapErr := c.mapper.From(*categories).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleGetCategoryByID(ctx echo.Context) error {

	uuid := ctx.Param("category_id")

	categories, err := c.categoryService.GetCategoryByUUID(uuid)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := viewmodel.Category{}
	mapErr := c.mapper.From(*categories).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleCreateCategory(ctx echo.Context) error {

	var input entity.Category

	jsonErr := ctx.Bind(&input)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		return ctx.JSON(err.StatusCode(), err)
	}

	category, err := c.categoryService.CreateCategory(input)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := viewmodel.Category{}
	mapErr := c.mapper.From(*category).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) handleUpdateCategoryByID(ctx echo.Context) error {

	uuid := ctx.Param("category_id")

	var category entity.Category
	jsonErr := ctx.Bind(&category)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		return ctx.JSON(err.StatusCode(), err)
	}

	err := c.categoryService.UpdateCategoryByID(uuid, category)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *Controller) handleDeleteCategoryByID(ctx echo.Context) error {

	uuid := ctx.Param("category_id")

	err := c.categoryService.DeleteCategoryByID(uuid)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

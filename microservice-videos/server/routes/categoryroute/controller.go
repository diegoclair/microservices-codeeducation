package categoryroute

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/entity"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/server/viewmodel"
	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/utils/routeutils"
	"github.com/gin-gonic/gin"
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

func (c *Controller) handleGetCategories(ctx *gin.Context) {

	categories, err := c.categoryService.GetCategories()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *Controller) handleGetCategoryByID(ctx *gin.Context) {

	id, err := routeutils.GetAndValidateIntParam(ctx, "category_id", false)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	categories, err := c.categoryService.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	response := viewmodel.Category{}
	mapErr := c.mapper.From(*categories).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do mapper: " + fmt.Sprint(mapErr))
		ctx.JSON(err.StatusCode(), err)
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleCreateCategory(ctx *gin.Context) {

	var input entity.Category

	jsonErr := ctx.ShouldBindJSON(&input)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		ctx.JSON(err.StatusCode(), err)
		return
	}

	category, err := c.categoryService.CreateCategory(input)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

func (c *Controller) handleUpdateCategoryByID(ctx *gin.Context) {

	id, err := routeutils.GetAndValidateIntParam(ctx, "category_id", false)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	var category entity.Category
	jsonErr := ctx.ShouldBindJSON(&category)
	if jsonErr != nil {
		err := resterrors.NewBadRequestError("Invalid json body")
		ctx.JSON(err.StatusCode(), err)
		return
	}

	err = c.categoryService.UpdateCategoryByID(id, category)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (c *Controller) handleDeleteCategoryByID(ctx *gin.Context) {

	id, err := routeutils.GetAndValidateIntParam(ctx, "category_id", false)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	err = c.categoryService.DeleteCategoryByID(id)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}

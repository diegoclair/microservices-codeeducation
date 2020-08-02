package categoryroute

import (
	"net/http"
	"sync"

	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
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
}

//NewController to handle requests
func NewController(categoryService contract.CategoryService) *Controller {
	once.Do(func() {
		instance = &Controller{
			categoryService: categoryService,
		}
	})
	return instance
}

func (s *Controller) handleGetCategories(c *gin.Context) {

	categories, err := s.categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (s *Controller) handleGetCategoryByID(c *gin.Context) {

	id, err := routeutils.GetAndValidateIntParam(c, "category_id", false)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	categories, err := s.categoryService.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

package categoryroute

import (
	"sync"

	"github.com/diegoclair/microservices-codeeducation/tree/master/microservice-videos/domain/contract"
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

// handlePing - handle a Ping request
func (s *Controller) handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

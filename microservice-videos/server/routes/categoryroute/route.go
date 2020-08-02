package categoryroute

import (
	"github.com/gin-gonic/gin"
)

// CategoriesRouter holds the category handlers
type CategoriesRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new CategoriesRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *CategoriesRouter {
	return &CategoriesRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *CategoriesRouter) RegisterRoutes() {
	r.router.GET("/api/categories", r.ctrl.handleGetCategories)
	r.router.GET("/api/categories/:category_id", r.ctrl.handleGetCategoryByID)
	r.router.POST("/api/categories/:category_id", r.ctrl.handleCreateCategory)
	r.router.PUT("/api/categories/:category_id", r.ctrl.handleUpdateCategoryByID)
}

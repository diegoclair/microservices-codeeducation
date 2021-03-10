package genreroute

import (
	"github.com/gin-gonic/gin"
)

// GenresRouter holds the genre handlers
type GenresRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new GenresRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *GenresRouter {
	return &GenresRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *GenresRouter) RegisterRoutes() {
	r.router.GET("/api/genres", r.ctrl.handleGetGenres)
	r.router.GET("/api/genres/:genre_id", r.ctrl.handleGetGenreByID)
	r.router.POST("/api/genres", r.ctrl.handleCreateGenre)
	r.router.PUT("/api/genres/:genre_id", r.ctrl.handleUpdateGenreByID)
	r.router.DELETE("/api/genres/:genre_id", r.ctrl.handleDeleteGenreByID)
}

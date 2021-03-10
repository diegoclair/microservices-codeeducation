package genreroute

import "github.com/labstack/echo/v4"

// GenresRouter holds the genre handlers
type GenresRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new GenresRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *GenresRouter {
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

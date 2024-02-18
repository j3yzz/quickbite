package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Router struct {
	webServer *echo.Echo
}

func New(webserver *echo.Echo) *Router {
	return &Router{webServer: webserver}
}

func (r *Router) Provide() {
	apiRouter := r.webServer.Group("/api")

	v1Router := apiRouter.Group("/v1")

	v1Router.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"msg": "Hello",
		})
	})
}

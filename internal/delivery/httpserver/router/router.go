package router

import (
	"github.com/j3yzz/quickbite/internal/delivery/httpserver/handler"
	"github.com/labstack/echo/v4"
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

	handler.Health{}.Register(v1Router)
}

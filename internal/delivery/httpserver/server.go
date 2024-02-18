package httpserver

import (
	"errors"
	"github.com/j3yzz/quickbite/internal/config"
	"github.com/j3yzz/quickbite/internal/delivery/httpserver/router"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server struct {
	config config.Server
}

func New(cfg config.Server) *Server {
	return &Server{config: cfg}
}

func (s *Server) Provide() {
	e := echo.New()

	r := router.New(e)
	r.Provide()

	if err := e.Start(":" + s.config.Port); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("echo init failed:", err)
	}
}

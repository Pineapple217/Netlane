package server

import (
	"github.com/Pineapple217/Netlane/pkg/handler"
	"github.com/Pineapple217/Netlane/pkg/static"
	"github.com/labstack/echo/v4"
)

func (server *Server) RegisterRoutes(hdlr *handler.Handler) {
	e := server.e

	s := e.Group("/static")
	s.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("Cache-Control", "public, max-age=31536000, immutable")
			return next(c)
		}
	})
	s.StaticFS("/", echo.MustSubFS(static.PublicFS, "public"))

	e.GET("/", hdlr.Home)

}

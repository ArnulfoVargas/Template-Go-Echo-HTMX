package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(template templ.Component, c echo.Context) error {
  return template.Render(c.Request().Context(), c.Response().Writer)
}

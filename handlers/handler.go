package handlers

import (
	"github.com/labstack/echo/v4"
)

func Handle(c *echo.Echo) {
  c.GET("/hello", hello)
}

func hello(c echo.Context) error {
  dummy := models.NewDummy()
  return render(templates.Hello(dummy.Name), c)
}

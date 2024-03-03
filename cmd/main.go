package main

import (
	"os"
	"strings"

	"github.com/ArnulfoVargas/Template-Go-Echo-HTMX/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
  godotenv.Load()
	app := echo.New()

  handlers.Handle(app)
  app.Static("/", "public")
  app.File("/", "views/index.html")

  panic(app.Start(getPort()))
}

func getPort() string {
  sb := strings.Builder{}
  
  sb.WriteByte(':')
  sb.WriteString(os.Getenv("PORT"))

  return sb.String()
}

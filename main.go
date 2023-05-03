package main

import (
	"penta/handler"
	"penta/service"

	"github.com/gin-gonic/gin"
)

func main() {
	dirService := service.NewDirListService()

	app := gin.New()

	handler.AddHandler(app, dirService)

	app.Run("0.0.0.0:8080")
}

package main

import (
	"github.com/IcaroSilvaFK/key-gen/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	app.Use(cors.Default())

	router := app.Group("/api/_v1")

	routes.ApplicationRoutes(router)

	app.Run()
}

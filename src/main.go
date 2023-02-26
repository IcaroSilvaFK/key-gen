package main

import (
	"github.com/IcaroSilvaFK/key-gen/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	router := app.Group("/api/_v1")

	router.Use(cors.Default())
	routes.ApplicationRoutes(router)

	app.Run()
}

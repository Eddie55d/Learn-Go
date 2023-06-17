package main

import (
	"ice-cream-app/internal/database"
	routehandlers "ice-cream-app/internal/handlers/route-handlers"

	logger "ice-cream-app/internal/handlers/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := database.ConnectDB()
	defer DB.Close()

	router := gin.Default()

	router.Use(logger.LoggerInfo)

	router.GET("/icecreams", routehandlers.GetIceCreams(DB))
	router.GET("/icecreams/:id", routehandlers.GetIceCreamByID(DB))
	router.POST("/icecreams", routehandlers.PostIceCream(DB))
	router.PUT("/icecreams/:id", routehandlers.PutIceCream(DB))
	router.DELETE("/icecreams/:id", routehandlers.DeleteIceCream(DB))

	router.Run("localhost:8080")
}

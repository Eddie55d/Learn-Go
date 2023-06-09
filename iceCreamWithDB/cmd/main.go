package main

import (
	routehandlers "ice-cream-app/internal/handlers/route-handlers"

	logger "ice-cream-app/internal/handlers/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(logger.LoggerInfo)

	router.GET("/icecreams", routehandlers.GetIceCreams)
	router.GET("/icecreams/:id", routehandlers.GetIceCreamByID)
	router.POST("/icecreams", routehandlers.PostIceCream)
	router.PUT("/icecreams/:id", routehandlers.PutIceCream)
	router.DELETE("/icecreams/:id", routehandlers.DeleteIceCream)

	router.Run("localhost:8080")

}

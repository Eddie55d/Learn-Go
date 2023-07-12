package router

import (
	deleteicecream "ice-cream-app/internal/controller/http/v1/delete-ice-cream"
	getallicecreams "ice-cream-app/internal/controller/http/v1/get-all-ice-creams"
	geticecreambyid "ice-cream-app/internal/controller/http/v1/get-ice-cream-by-id"
	posticecream "ice-cream-app/internal/controller/http/v1/post-ice-cream"
	puticecream "ice-cream-app/internal/controller/http/v1/put-ice-cream"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	GetAllIceCreamHandler  getallicecreams.GetAllIceCreamHandler
	GetIceCreamByIDHandler geticecreambyid.IceCreamByIdHandler
	DeleteIceCreamHandler  deleteicecream.DeleteIceCreamHandler
	PostIceCreamHandler    posticecream.PostIceCreamHandler
	PutIceCreamHandler     puticecream.PutIceCreamHandler
}

func New(h Handlers, middleware ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middleware...)
	//router.Use(logger.LoggerInfo)

	router.GET("/icecreams", h.GetAllIceCreamHandler.GetIceCreams)
	router.GET("/icecreams/:id", h.GetIceCreamByIDHandler.GetIceCreamByID)
	router.DELETE("/icecreams/:id", h.DeleteIceCreamHandler.DelIceCream)
	router.POST("/icecreams", h.PostIceCreamHandler.PostIceCream)
	router.PUT("/icecreams/:id", h.PutIceCreamHandler.PostIceCream)
	return router
}

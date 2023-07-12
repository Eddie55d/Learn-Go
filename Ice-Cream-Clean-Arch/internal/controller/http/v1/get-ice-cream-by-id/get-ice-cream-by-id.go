package geticecreambyid

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"ice-cream-app/internal/controller/services"
	"ice-cream-app/internal/models"
)

type IceCreamByIdHandler struct {
	service services.IceCreamService
}

func NewIceCreamHandler(s services.IceCreamService) *IceCreamByIdHandler {
	return &IceCreamByIdHandler{service: s}
}

func (h *IceCreamByIdHandler) GetIceCreamByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id field is empty"})
		return
	}

	iceCream, err := h.service.GetIceCreamByID(id)

	switch {
	case err != nil && !errors.Is(err, models.ErrSqlNoRow):
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	case errors.Is(err, models.ErrSqlNoRow):
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
		logrus.Info("ice cream not found")
		return
	default:
		c.IndentedJSON(http.StatusOK, iceCream)
		return
	}
}

package deleteicecream

import (
	"errors"
	"fmt"
	"ice-cream-app/internal/controller/services"
	"ice-cream-app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DeleteIceCreamHandler struct {
	service services.IceCreamService
}

func NewIceCreamHandler(s services.IceCreamService) *DeleteIceCreamHandler {
	return &DeleteIceCreamHandler{service: s}
}

func (h *DeleteIceCreamHandler) DelIceCream(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": `field "id" is empty`})
		return
	}

	err := h.service.DeleteIceCream(id)

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
		deleteIceCreamMsg := fmt.Sprintf("ice cream with id %s is removed", id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": deleteIceCreamMsg})
		logrus.Info(deleteIceCreamMsg)
	}
}

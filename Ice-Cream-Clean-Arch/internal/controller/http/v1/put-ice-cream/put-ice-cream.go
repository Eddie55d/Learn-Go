package puticecream

import (
	"errors"
	"fmt"
	"ice-cream-app/internal/controller/services"
	"ice-cream-app/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PutIceCreamHandler struct {
	service services.IceCreamService
}

func NewIceCreamHandler(s services.IceCreamService) *PutIceCreamHandler {
	return &PutIceCreamHandler{service: s}
}

func (h *PutIceCreamHandler) PostIceCream(c *gin.Context) {

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
	}

	var updateIceCream models.IceCreamPost
	if err := c.BindJSON(&updateIceCream); err != nil {
		logrus.Warn(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are not filled correctly"})
		return
	}

	validUpdate, err := updateIceCream.ValidUpdate(&iceCream)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		logrus.Error(err)
		return
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.service.PutIceCream(paramID, validUpdate)
	switch {
	case err != nil && !errors.Is(err, models.ErrSqlNoRow):
		logrus.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	case errors.Is(err, models.ErrSqlNoRow) || res == 0:
		c.JSON(http.StatusBadRequest, gin.H{"message": "nothing has been updated"})
		logrus.Error("nothing has been updated")
		return
	default:
		var updateMsg = fmt.Sprintf("ice cream with id %d updated!", paramID)
		c.IndentedJSON(http.StatusCreated, gin.H{"message": updateMsg})
		logrus.Info(updateMsg)
	}

}

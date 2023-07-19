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
		logrus.Info("ice cream not found")
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
		return
	}

	var updateIceCream models.IceCreamPost
	if err := c.BindJSON(&updateIceCream); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are not filled correctly"})
		return
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.service.PutIceCream(paramID, &iceCream, &updateIceCream)
	switch {
	case
		errors.Is(err, models.ErrTitleEmpty) ||
			errors.Is(err, models.ErrCompositionEmpty) ||
			errors.Is(err, models.ErrDateOfManufacture) ||
			errors.Is(err, models.ErrExpirationDate) ||
			errors.Is(err, models.ErrPriceEmpty) ||
			errors.Is(err, models.ErrPrice):
		logrus.Error(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case err != nil && !errors.Is(err, models.ErrSqlNoRow):
		logrus.Error(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	case errors.Is(err, models.ErrSqlNoRow):
		logrus.Error("nothing has been updated")
		c.JSON(http.StatusBadRequest, gin.H{"message": "nothing has been updated"})
		return
	default:
		var updateMsg = fmt.Sprintf("ice cream with id %d updated!", paramID)
		logrus.Info(updateMsg)
		c.IndentedJSON(http.StatusCreated, gin.H{"message": updateMsg})
	}

}

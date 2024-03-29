package posticecream

import (
	"errors"
	"fmt"
	"ice-cream-app/internal/controller/services"
	"ice-cream-app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostIceCreamHandler struct {
	service services.IceCreamService
}

func NewIceCreamHandler(s services.IceCreamService) *PostIceCreamHandler {
	return &PostIceCreamHandler{service: s}
}

func (h *PostIceCreamHandler) PostIceCream(c *gin.Context) {
	var newIceCream models.IceCreamPost

	if err := c.BindJSON(&newIceCream); err != nil {
		logrus.Warn(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are not filled correctly"})
		return
	}

	res, err := h.service.PostIceCream(newIceCream)
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
		logrus.Error("nothing added")
		c.JSON(http.StatusBadRequest, gin.H{"message": "nothing added"})
		return
	default:
		var createdMsg = fmt.Sprintf("ice cream with id %d created!", res)
		logrus.Info(createdMsg)
		c.IndentedJSON(http.StatusCreated, gin.H{"message": createdMsg})
	}

}

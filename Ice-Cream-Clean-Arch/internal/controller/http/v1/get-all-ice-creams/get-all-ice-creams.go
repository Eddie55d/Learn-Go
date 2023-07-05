package getallicecreams

import (
	"ice-cream-app/internal/controller/services"
	"ice-cream-app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetAllIceCreamHandler struct {
	service services.IceCreamService
}

func NewIceCreamHandler(s services.IceCreamService) *GetAllIceCreamHandler {
	return &GetAllIceCreamHandler{service: s}
}

func (h *GetAllIceCreamHandler) GetIceCreams(c *gin.Context) {

	limit := c.Query("limit")
	offset := c.Query("offset")

	_, isManufactureDate := c.GetQuery("date_of_manufacture")
	_, isExpirationDate := c.GetQuery("expiration_date")
	_, isPrice := c.GetQuery("price")

	qparams := models.QueryParams{
		Limit:           limit,
		Offset:          offset,
		ManufactureDate: isManufactureDate,
		ExpirationDate:  isExpirationDate,
		Price:           isPrice,
	}

	iceCreams, err := h.service.GetIceCreams(qparams)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unable to get all ice creams"})
		logrus.Error("Unable to get all ice creams")
		return
	}

	if len(iceCreams) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream list is empty"})
		logrus.Info("ice cream list is empty")
		return
	}
	c.IndentedJSON(http.StatusOK, iceCreams)
}

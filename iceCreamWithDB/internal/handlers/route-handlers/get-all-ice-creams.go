package routehandlers

import (
	"fmt"
	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetIceCreams(c *gin.Context) {

	db := database.ConnectDB()
	defer db.Close()

	selectRow := "SELECT icecream_id, title, composition, date_of_manufacture, expiration_date, price FROM icecreams WHERE is_deleted IS NULL ORDER BY icecream_id"

	_, isManufactureDate := c.GetQuery("date_of_manufacture")
	_, isExpirationDate := c.GetQuery("expiration_date")
	_, isPrice := c.GetQuery("price")

	switch {
	case isManufactureDate:
		selectRow = "SELECT icecream_id, title, composition, date_of_manufacture, expiration_date, price FROM icecreams WHERE is_deleted IS NULL ORDER BY date_of_manufacture "
	case isExpirationDate:
		selectRow = "SELECT icecream_id, title, composition, date_of_manufacture, expiration_date, price FROM icecreams WHERE is_deleted IS NULL ORDER BY expiration_date "
	case isPrice:
		selectRow = "SELECT icecream_id, title, composition, date_of_manufacture, expiration_date, price FROM icecreams WHERE is_deleted IS NULL ORDER BY price "
	}

	limit := c.Query("limit")
	if limit != "" {
		selectRow += fmt.Sprintf(` LIMIT %s`, limit)
	}

	offset := c.Query("offset")
	if offset != "" {
		selectRow += fmt.Sprintf(` OFFSET %s`, offset)
	}

	rows, err := db.Query(selectRow)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer rows.Close()

	iceCreams := []models.IceCream{}

	for rows.Next() {
		p := models.IceCream{}
		err := rows.Scan(&p.Icecream_id, &p.Title, &p.Сomposition, &p.DateOfManufacture, &p.ExpirationDate, &p.Price)
		if err != nil {
			logrus.Error(err)
			return
		}
		iceCreams = append(iceCreams, p)
	}

	if len(iceCreams) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream list is empty"})
		logrus.Info("ice cream list is empty")
		return
	}
	c.IndentedJSON(http.StatusOK, iceCreams)

}

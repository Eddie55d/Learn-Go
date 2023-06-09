package routehandlers

import (
	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetIceCreams(c *gin.Context) {

	db := database.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM icecreams")
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

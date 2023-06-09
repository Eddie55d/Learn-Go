package routehandlers

import (
	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostIceCream(c *gin.Context) {
	var newIceCream models.IceCreamPost

	db := database.ConnectDB()
	defer db.Close()

	if err := c.BindJSON(&newIceCream); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Fields are not filled correctly"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO icecreams (title, composition, date_of_manufacture, expiration_date, price) VALUES ($1, $2, $3, $4, $5) RETURNING icecream_id")
	if err != nil {
		logrus.Error(err)
		return
	}

	defer stmt.Close()

	validIce, msg, err := newIceCream.ValidPost()
	if err != nil && msg == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		logrus.Error(err)
		return
	}

	if msg != "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": msg})
		logrus.Errorf(msg)
		return
	}

	res, err := stmt.Exec(validIce.Title, validIce.Сomposition, validIce.DateOfManufacture, validIce.ExpirationDate, validIce.Price)
	if err != nil {
		logrus.Error(err)
	}

	lastId, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
	}

	if lastId == 0 {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "nothing added"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "ice cream created!"})
}

package routehandlers

import (
	"database/sql"

	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PutIceCream(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()

	var iceCream models.IceCream
	var upIce models.IceCreamPost
	id := c.Param("id")

	errGet := db.QueryRow("SELECT * FROM icecreams WHERE icecream_id = $1", id).
		Scan(&iceCream.Icecream_id, &iceCream.Title, &iceCream.Сomposition, &iceCream.DateOfManufacture, &iceCream.ExpirationDate, &iceCream.Price)

	if errGet == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
	}

	if err := c.BindJSON(&upIce); err != nil {
		logrus.Error(err)
		return
	}

	stmt, err := db.Prepare("UPDATE icecreams SET title=$1, composition=$2, date_of_manufacture=$3, expiration_date=$4, price=$5 WHERE icecream_id=$6")
	if err != nil {
		logrus.Error(err)
		return
	}

	defer stmt.Close()

	validUpdate, msg, err := upIce.ValidUpdate(&iceCream)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": msg, "error": err})
		logrus.Error(err)
		return
	}

	res, err := stmt.Exec(validUpdate.Title, validUpdate.Сomposition, validUpdate.DateOfManufacture, validUpdate.ExpirationDate, validUpdate.Price, id)
	if err != nil {
		logrus.Error(err)
		return
	}

	n, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return
	}

	if n == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "ice cream updated"})

}

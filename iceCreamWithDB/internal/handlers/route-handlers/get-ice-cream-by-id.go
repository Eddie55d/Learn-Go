package routehandlers

import (
	"database/sql"

	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetIceCreamByID(c *gin.Context) {

	var iceCream models.IceCream

	db := database.ConnectDB()
	defer db.Close()

	id := c.Param("id")

	err := db.QueryRow("SELECT * FROM icecreams WHERE icecream_id = $1", id).
		Scan(&iceCream.Icecream_id, &iceCream.Title, &iceCream.Сomposition, &iceCream.DateOfManufacture, &iceCream.ExpirationDate, &iceCream.Price)

	switch {
	case err != nil && err != sql.ErrNoRows:
		logrus.Error(err)
		return
	case err == sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
		logrus.Info("ice cream not found")
		return
	default:
		c.IndentedJSON(http.StatusOK, iceCream)
		return
	}
}

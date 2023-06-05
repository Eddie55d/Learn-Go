package routehandlers

import (
	"fmt"
	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIceCreams(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM icecreams")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	iceCreams := []models.IceCream{}

	for rows.Next() {
		p := models.IceCream{}
		err := rows.Scan(&p.Icecream_id, &p.Title, &p.Сomposition, &p.DateOfManufacture, &p.ExpirationDate, &p.Price)
		if err != nil {
			fmt.Println(err)
			//log.Fatal(err)
			continue
		}
		iceCreams = append(iceCreams, p)
	}

	if len(iceCreams) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream list is empty"})
	}
	c.IndentedJSON(http.StatusOK, iceCreams)
}

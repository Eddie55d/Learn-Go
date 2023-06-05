package routehandlers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func PutIceCream(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()

	var iceCream models.IceCream
	var upIce models.IceCreamUp
	id := c.Param("id")

	errGet := db.QueryRow("SELECT * FROM icecreams WHERE icecream_id = $1", id).
		Scan(&iceCream.Icecream_id, &iceCream.Title, &iceCream.Сomposition, &iceCream.DateOfManufacture, &iceCream.ExpirationDate, &iceCream.Price)

	if errGet == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
	}

	if err := c.BindJSON(&upIce); err != nil {
		fmt.Println(err)
		return
	}

	stmt, err := db.Prepare("UPDATE icecreams SET title=$1, composition=$2, date_of_manufacture=$3, expiration_date=$4, price=$5 WHERE icecream_id=$6")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	if upIce.Title == "" {
		upIce.Title = iceCream.Title
	}

	if upIce.Сomposition == "" {
		upIce.Сomposition = iceCream.Сomposition
	}

	if upIce.DateOfManufacture == "" {
		upIce.DateOfManufacture = iceCream.DateOfManufacture.Format("2006-01-02")
	}

	if upIce.ExpirationDate == "" {
		upIce.ExpirationDate = iceCream.ExpirationDate.Format("2006-01-02")
	}

	if upIce.Price == "" {
		upIce.Price = iceCream.Price
	}

	dtST, err := time.Parse("2006-01-02", upIce.DateOfManufacture)
	if err != nil {
		fmt.Println(err)
	}

	dtEnd, err := time.Parse("2006-01-02", upIce.ExpirationDate)
	if err != nil {
		fmt.Println(err)
	}

	pr, err := strconv.ParseFloat(upIce.Price, 32)
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(upIce.Title, upIce.Сomposition, dtST, dtEnd, pr, id)
	if err != nil {
		fmt.Println(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	if n == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "ice cream updated"})

}

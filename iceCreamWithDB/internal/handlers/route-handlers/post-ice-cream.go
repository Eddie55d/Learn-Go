package routehandlers

import (
	"fmt"
	"ice-cream-app/internal/database"
	"ice-cream-app/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PostIceCream(c *gin.Context) {
	var newIceCream models.IceCreamPost

	db := database.ConnectDB()
	defer db.Close()

	if err := c.BindJSON(&newIceCream); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Fields are not filled correctly"})
		return
	}

	fmt.Println(newIceCream)

	if newIceCream.Title == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "title field is empty"})
		return
	}

	if newIceCream.Сomposition == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "composition field is empty"})
		return
	}

	if newIceCream.Price == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "price field is empty"})
		return
	}

	tm, err := time.Parse("2006-01-02", newIceCream.DateOfManufacture)
	if err != nil {
		fmt.Println(err)
	}

	if tm.IsZero() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "DateOfManufacture field is empty"})
		return
	}

	tExp, err := time.Parse("2006-01-02", newIceCream.DateOfManufacture)
	if err != nil {
		fmt.Println(err)
	}

	if tExp.IsZero() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "ExpirationDate field is empty"})
		return
	}

	dtST, err := time.Parse("2006-05-01", newIceCream.DateOfManufacture)
	if err != nil {
		fmt.Println(err)
	}

	dtEnd, err := time.Parse("2006-05-01", newIceCream.ExpirationDate)
	if err != nil {
		fmt.Println(err)
	}

	pr, err := strconv.ParseFloat(newIceCream.Price, 32)
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO icecreams (title, composition, date_of_manufacture, expiration_date, price) VALUES ($1, $2, $3, $4, $5) RETURNING icecream_id")
	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(newIceCream.Title, newIceCream.Сomposition, dtST, dtEnd, pr)
	if err != nil {
		fmt.Println(err)
	}

	lastId, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	if lastId == 0 {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "nothing added"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "ice cream created!"})
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type iceCream struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Сomposition       string  `json:"composition"`
	DateOfManufacture string  `json:"date-start"`
	ExpirationDate    string  `json:"date-end"`
	Price             float64 `json:"price"`
}

type iceCreamUpdate struct {
	ID                string  `json:"id,omitempty"`
	Title             string  `json:"title,omitempty"`
	Сomposition       string  `json:"composition,omitempty"`
	DateOfManufacture string  `json:"date-start,omitempty"`
	ExpirationDate    string  `json:"date-end,omitempty"`
	Price             float64 `json:"price,omitempty"`
}

var iceCreams = []iceCream{
	{
		ID:                "1",
		Title:             "name1",
		Сomposition:       "dhiuhihd, i0wis, wjsjw",
		DateOfManufacture: "12.12.2009",
		ExpirationDate:    "12.12.2018",
		Price:             9.00,
	},
	{
		ID:                "2",
		Title:             "name2",
		Сomposition:       "hdis,jidjoj,jddolss,",
		DateOfManufacture: "21.09.2021",
		ExpirationDate:    "22.09.2022",
		Price:             15.67,
	},
	{
		ID:                "3",
		Title:             "name3",
		Сomposition:       "hdis,jidjoj,jddolss",
		DateOfManufacture: "21.09.2021",
		ExpirationDate:    "22.09.2022",
		Price:             89.67,
	},
}

func getIceCreams(c *gin.Context) {
	if len(iceCreams) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream list is empty"})
	}
	c.IndentedJSON(http.StatusOK, iceCreams)
}

func deleteIceCream(c *gin.Context) {
	id := c.Param("id")

	for idx, ice := range iceCreams {
		if ice.ID == id {

			iceCreams = append(iceCreams[:idx], iceCreams[idx+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "ice cream deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
}

func getIceCreamByID(c *gin.Context) {
	id := c.Param("id")

	for _, ice := range iceCreams {
		if ice.ID == id {
			c.IndentedJSON(http.StatusOK, ice)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
}

func postIceCream(c *gin.Context) {
	var newIceCream iceCream

	if err := c.BindJSON(&newIceCream); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Fields are not filled correctly"})
		return
	}

	iceCreams = append(iceCreams, newIceCream)
	c.IndentedJSON(http.StatusCreated, newIceCream)
}

func putIceCreams(c *gin.Context) {
	id := c.Param("id")

	var updateIceCream iceCreamUpdate
	for idx, ice := range iceCreams {
		if ice.ID == id {
			oldIce := ice
			iceCreams = append(iceCreams[:idx], iceCreams[idx+1:]...)
			c.BindJSON(&updateIceCream)

			updateIceCream = iceCreamUpdate{
				ID:                oldIce.ID,
				Title:             updateIceCream.Title,
				Сomposition:       updateIceCream.Сomposition,
				DateOfManufacture: updateIceCream.DateOfManufacture,
				ExpirationDate:    updateIceCream.ExpirationDate,
				Price:             updateIceCream.Price,
			}

			if updateIceCream.Title == "" {
				updateIceCream.Сomposition = oldIce.Сomposition
			} else if updateIceCream.Сomposition == "" {
				updateIceCream.Title = oldIce.Title
			} else if updateIceCream.DateOfManufacture == "" {
				updateIceCream.DateOfManufacture = oldIce.DateOfManufacture
			} else if updateIceCream.ExpirationDate == "" {
				updateIceCream.ExpirationDate = oldIce.ExpirationDate
			} else if updateIceCream.Price == 0 {
				updateIceCream.Price = oldIce.Price
			}

			iceCreams = append(iceCreams, iceCream(updateIceCream))
			c.IndentedJSON(http.StatusCreated, gin.H{"message": "ice cream updated"})
			return
		}

	}
	c.JSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
}

func main() {

	router := gin.Default()

	router.GET("/icecreams", getIceCreams)
	router.GET("/icecreams/:id", getIceCreamByID)
	router.POST("/icecreams", postIceCream)
	router.PUT("/icecreams/:id", putIceCreams)
	router.DELETE("/icecreams/:id", deleteIceCream)

	router.Run("localhost:8080")

}

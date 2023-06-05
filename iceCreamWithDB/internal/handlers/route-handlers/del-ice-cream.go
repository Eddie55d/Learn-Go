package routehandlers

import (
	"fmt"
	"log"

	"ice-cream-app/internal/database"

	"github.com/gin-gonic/gin"

	"net/http"
)

func DeleteIceCream(c *gin.Context) {

	db := database.ConnectDB()
	defer db.Close()

	id := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM icecreams WHERE icecream_id=$1")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	if n == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ice cream deleted"})

}

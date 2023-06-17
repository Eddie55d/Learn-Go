package routehandlers

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"net/http"
)

func DeleteIceCream(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		id := c.Param("id")

		stmt, err := db.Prepare("UPDATE icecreams SET is_deleted = $2 WHERE icecream_id = $1 AND is_deleted IS NULL")
		if err != nil {
			logrus.Error(err)
			return
		}

		defer stmt.Close()

		deleteTime := time.Now()

		res, err := stmt.Exec(id, deleteTime)
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
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ice cream not found"})
			logrus.Info("ice cream not found")
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ice cream deleted"})
		logrus.Info("ice cream deleted")
	}
}

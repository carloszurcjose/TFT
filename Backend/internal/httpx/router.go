package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		//Optioanlly ping DB to confirm connectivity
		if err := db.Exec("SELECT 1").Error; err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"ok": false, "db": "down", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "db": "up"})
	})
	return r
}

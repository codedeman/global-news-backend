package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticlesHandler(c *gin.Context) {
	articles := GetArticles()
	c.JSON(http.StatusOK, gin.H{
		"count": len(articles),
		"data":  articles,
	})
}

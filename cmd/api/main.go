package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yourname/global-news-backend/internal/article"
	"github.com/yourname/global-news-backend/internal/rss"
)

func main() {

	// Fetch RSS feeds khi app start
	go rss.ProcessFeeds()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Global News Backend Running 🚀",
		})
	})

	router.GET("/articles", article.GetArticlesHandler)

	log.Println("Server running on :8080")

	router.Run(":8080")
}

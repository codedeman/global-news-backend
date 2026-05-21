package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/codedeman/global-news-backend/internal/article"
	"github.com/codedeman/global-news-backend/internal/rss"
)

func main() {

	go rss.StartWorker()

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
	router.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/articles", article.GetArticlesHandler)

	log.Println("Server running on :8080")

	router.Run(":8080")
}

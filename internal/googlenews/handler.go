package googlenews

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/codedeman/neura-backend/internal/rss"
)

func Handler(c *gin.Context) {
	country := c.Query("country")
	if country == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing ?country= parameter (e.g. VN, SG, US)"})
		return
	}

	feedURL := rss.GoogleNewsFeedURL(country)
	if feedURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported country code: " + country})
		return
	}

	data, err := rss.Fetch(feedURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch feed: " + err.Error()})
		return
	}

	articles, err := rss.Parse(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse feed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"country": country,
		"count":   len(articles),
		"data":    articles,
	})
}

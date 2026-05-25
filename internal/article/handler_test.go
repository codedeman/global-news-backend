package article

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGetArticlesHandler_EmptyStore(t *testing.T) {
	SetArticles([]Article{})
	defer SetArticles([]Article{})

	router := gin.New()
	router.GET("/articles", GetArticlesHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/articles", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	assert.Equal(t, float64(0), body["count"])
}

func TestGetArticlesHandler_WithArticles(t *testing.T) {
	articles := []Article{
		{Title: "Test 1", Link: "https://example.com/1"},
		{Title: "Test 2", Link: "https://example.com/2"},
	}
	SetArticles(articles)
	defer SetArticles([]Article{})

	router := gin.New()
	router.GET("/articles", GetArticlesHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/articles", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	assert.Equal(t, float64(2), body["count"])

	data := body["data"].([]any)
	assert.Len(t, data, 2)
}

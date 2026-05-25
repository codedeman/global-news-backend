package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGetArticles(t *testing.T) {
	defer SetArticles([]Article{})

	input := []Article{
		{Title: "A", Link: "https://example.com/a"},
		{Title: "B", Link: "https://example.com/b"},
	}

	SetArticles(input)
	result := GetArticles()

	assert.Len(t, result, 2)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[1].Title)
}

func TestGetArticles_EmptyStore(t *testing.T) {
	SetArticles([]Article{})

	result := GetArticles()

	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestSetArticles_OverwritesPrevious(t *testing.T) {
	defer SetArticles([]Article{})

	SetArticles([]Article{{Title: "First"}})
	SetArticles([]Article{{Title: "Second"}, {Title: "Third"}})

	result := GetArticles()

	assert.Len(t, result, 2)
	assert.Equal(t, "Second", result[0].Title)
}

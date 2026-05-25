package rss

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var validRSSFeed = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Test Feed</title>
    <item>
      <title>Article One</title>
      <description>Description one</description>
      <link>https://example.com/1</link>
      <pubDate>Mon, 01 Jan 2024 10:00:00 +0000</pubDate>
    </item>
    <item>
      <title>Article Two</title>
      <description>Description two</description>
      <link>https://example.com/2</link>
      <pubDate>Tue, 02 Jan 2024 10:00:00 +0000</pubDate>
    </item>
  </channel>
</rss>`)

var rssWithNoDate = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Test Feed</title>
    <item>
      <title>No Date Article</title>
      <description>No date</description>
      <link>https://example.com/nodate</link>
    </item>
  </channel>
</rss>`)

var emptyRSSFeed = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Empty Feed</title>
  </channel>
</rss>`)

func TestParse_ValidFeed(t *testing.T) {
	articles, err := Parse(validRSSFeed)

	require.NoError(t, err)
	require.Len(t, articles, 2)

	assert.Equal(t, "Article One", articles[0].Title)
	assert.Equal(t, "Description one", articles[0].Description)
	assert.Equal(t, "https://example.com/1", articles[0].Link)

	assert.Equal(t, "Article Two", articles[1].Title)
	assert.Equal(t, "https://example.com/2", articles[1].Link)
}

func TestParse_InvalidXML(t *testing.T) {
	_, err := Parse([]byte("this is not xml"))

	assert.Error(t, err)
}

func TestParse_ItemWithPublishedDate(t *testing.T) {
	articles, err := Parse(validRSSFeed)

	require.NoError(t, err)
	require.Len(t, articles, 2)

	expected := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, articles[0].PublishedAt)
}

func TestParse_ItemWithoutPublishedDate(t *testing.T) {
	before := time.Now()
	articles, err := Parse(rssWithNoDate)
	after := time.Now()

	require.NoError(t, err)
	require.Len(t, articles, 1)

	assert.False(t, articles[0].PublishedAt.IsZero())
	assert.True(t, articles[0].PublishedAt.After(before) || articles[0].PublishedAt.Equal(before))
	assert.True(t, articles[0].PublishedAt.Before(after) || articles[0].PublishedAt.Equal(after))
}

func TestParse_EmptyFeed(t *testing.T) {
	articles, err := Parse(emptyRSSFeed)

	require.NoError(t, err)
	assert.Empty(t, articles)
}

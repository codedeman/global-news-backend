package rss

import (
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/codedeman/global-news-backend/internal/article"
)

func Parse(feedData []byte) ([]article.Article, error) {
	parser := gofeed.NewParser()

	feed, err := parser.ParseString(string(feedData))
	if err != nil {
		return nil, err
	}

	var articles []article.Article

	for _, item := range feed.Items {

		publishedAt := time.Now()

		if item.PublishedParsed != nil {
			publishedAt = *item.PublishedParsed
		}

		articles = append(articles, article.Article{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PublishedAt: publishedAt,
		})
	}

	return articles, nil
}

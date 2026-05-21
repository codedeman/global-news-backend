package rss

import (
	"log"
	"time"

	"github.com/codedeman/global-news-backend/internal/article"
)

func StartWorker() {
	ProcessFeeds()
	ticker := time.NewTicker(15 * time.Minute)
	for range ticker.C {
		ProcessFeeds()
	}
}

func ProcessFeeds() {
	all := []article.Article{}

	for _, source := range FeedSources {
		log.Println("Fetching:", source.Name)

		data, err := Fetch(source.URL)
		if err != nil {
			log.Println(err)
			continue
		}

		articles, err := Parse(data)
		if err != nil {
			log.Println(err)
			continue
		}

		all = append(all, articles...)
	}

	article.SetArticles(all)
	log.Printf("Fetched %d articles total", len(all))
}

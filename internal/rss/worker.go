package rss

import (
	"log"

	"github.com/codedeman/global-news-backend/internal/article"
)

func ProcessFeeds() {
	var all []article.Article

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

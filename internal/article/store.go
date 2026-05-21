package article

import "sync"

var (
	mu    sync.RWMutex
	store = []Article{}
)

func SetArticles(a []Article) {
	mu.Lock()
	defer mu.Unlock()
	store = a
}

func GetArticles() []Article {
	mu.RLock()
	defer mu.RUnlock()
	return store
}

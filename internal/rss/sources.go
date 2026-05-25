package rss

type FeedSource struct {
	Name string
	URL  string
}

var FeedSources = []FeedSource{
	{
		Name: "BBC",
		URL:  "https://feeds.bbci.co.uk/news/rss.xml",
	},
	{
		Name: "VNExpress",
		URL:  "https://vnexpress.net/rss/tin-moi-nhat.rss",
	},
	{
		Name: "7News Australia",
		URL:  "https://7news.com.au/feed",
	},
	{
		Name: "Vietnam News",
		URL:  "https://vietnamnews.vn/rss",
	},
	{
		Name: "VietnamPlus",
		URL:  "https://en.vietnamplus.vn/rss.html",
	},
	{
		Name: "VNA",
		URL:  "https://vnanet.vn/en/rss/",
	},
	{
		Name: "Google News VN",
		URL:  "https://news.google.com/rss?hl=vi&gl=VN&ceid=VN:vi",
	},
}

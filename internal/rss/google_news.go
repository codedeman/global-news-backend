package rss

import "fmt"

type countryLocale struct {
	lang string
	gl   string
}

var googleNewsCountries = map[string]countryLocale{
	"VN": {lang: "vi", gl: "VN"},
	"SG": {lang: "en", gl: "SG"},
	"US": {lang: "en", gl: "US"},
	"GB": {lang: "en", gl: "GB"},
	"AU": {lang: "en", gl: "AU"},
	"JP": {lang: "ja", gl: "JP"},
	"KR": {lang: "ko", gl: "KR"},
	"FR": {lang: "fr", gl: "FR"},
	"DE": {lang: "de", gl: "DE"},
	"TH":  {lang: "th", gl: "TH"},
	"ID":  {lang: "id", gl: "ID"},
	"KH":  {lang: "km", gl: "KH"},
	"CAM": {lang: "km", gl: "KH"},
	"LA":  {lang: "lo", gl: "LA"},
	"LAO": {lang: "lo", gl: "LA"},
}

// GoogleNewsFeedURL builds the Google News RSS URL for a given country code.
// Returns empty string if country code is not supported.
func GoogleNewsFeedURL(countryCode string) string {
	locale, ok := googleNewsCountries[countryCode]
	if !ok {
		return ""
	}
	return fmt.Sprintf("https://news.google.com/rss?hl=%s&gl=%s&ceid=%s:%s",
		locale.lang, locale.gl, locale.gl, locale.lang)
}

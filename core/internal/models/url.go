package models

import "net/url"

func stringToUrl(s string) url.URL {
	url, _ := url.Parse(s)
	return *url
}

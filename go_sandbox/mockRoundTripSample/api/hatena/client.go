package hatena

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	blogURL = "https://ema-hiro.hatenablog.com/feed"
	rssURL  = "https://ema-hiro.hatenablog.com/rss"
)

// FetchFeed fetches hatena blog feed.
func FetchFeed() (*Feed, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, blogURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("faield to request to hateba blog url. status: %v", resp.StatusCode)
	}

	feed := &Feed{}
	if err := xml.NewDecoder(resp.Body).Decode(feed); err != nil {
		return nil, err
	}
	return feed, nil
}

// FetchRSS fetches hatena blog rss feed.
func FetchRSS() (*RSS, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, rssURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("faield to request to hateba blog url. status: %v", resp.StatusCode)
	}

	rss := &RSS{}
	if err := xml.NewDecoder(resp.Body).Decode(rss); err != nil {
		return nil, err
	}
	return rss, nil
}

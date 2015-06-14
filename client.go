// Package bing is a Bing Search API client (only web implemented)
package bing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client provides a client to the Bing Search API
type Client struct {
	HTTP *http.Client
	// AuthKey is your API Auth Key from the Azure Data Marketplace
	AuthKey string
}

func (b *Client) prepareQuery(query string) string {
	path := &url.URL{Path: query}
	escaped := strings.Replace(path.String(), "%27", "'", -1)
	return fmt.Sprint("%27", escaped, "%27")
}

// Search queries the Web search API
func (b *Client) Search(query string) ([]Result, error) {
	u := fmt.Sprintf(
		"https://api.datamarket.azure.com/Bing/Search/v1/Web?$format=json&Query=%s",
		b.prepareQuery(query))

	req, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(b.AuthKey, b.AuthKey)

	resp, err := b.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, body)
	}

	var js struct {
		Data struct {
			Results []Result `json:"results"`
		} `json:"d"`
	}

	if err := json.Unmarshal(body, &js); err != nil {
		return nil, err
	}

	return js.Data.Results, nil
}

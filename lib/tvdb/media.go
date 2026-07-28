package tvdb

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getSeries(id string) (string, error) {
	req, err := c.AuthenticatedHttpRequest("GET", BaseURL+"/series/"+id+"/episodes/default/eng", nil)

	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil { 
		return "", err
	}

	return string(body), nil
}

// func (c *Client) getMovie(id string) (string, error) {

// }

func (c *Client) GetMedia(id string, kind string) (string, error) {
	if kind == "series" {
		return c.getSeries(id)
	}

	// if kind == "movie" {
	// 	return c.getMovie(id)
	// }

	return "", fmt.Errorf(`wrong "kind" value passed`)
}

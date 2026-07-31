package tvdb

import (
	"anime-meta/lib/core/fetch"
	"fmt"
)

func (c *Client) getSeries(id string) (any, error) {
	resp, err := c.TvdbAuthenticatedFetch("GET", BaseURL+"/series/"+id+"/episodes/default/eng", nil)
	if err != nil {
		return "", err
	}

	var seriesResp map[string]any

	fetch.ExtractResponseJsonBody(resp, &seriesResp)
	if err != nil {
		return "", err
	}

	return seriesResp, nil
}

// func (c *Client) getMovie(id string) (string, error) {

// }

func (c *Client) GetMedia(id string, kind string) (any, error) {
	if kind == "series" {
		return c.getSeries(id)
	}

	// if kind == "movie" {
	// 	return c.getMovie(id)
	// }

	return nil, fmt.Errorf(`wrong "kind" value passed`)
}

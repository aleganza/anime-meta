package tvdb

import (
	"anime-meta/lib/core/fetch"
)

func (c *Client) FetchSeries(id string) (TvdbSeriesResponse, error) {
	resp, err := c.TvdbAuthenticatedFetch("GET", BaseURL+"/series/"+id+"/episodes/default/eng", nil)
	if err != nil {
		return TvdbSeriesResponse{}, err
	}

	var seriesResp TvdbSeriesResponse

	fetch.ExtractResponseJsonBody(resp, &seriesResp)
	if err != nil {
		return TvdbSeriesResponse{}, err
	}

	return seriesResp, nil
}

func (c *Client) FetchMovie(id string) (TvdbMovieResponse, error) {
	resp, err := c.TvdbAuthenticatedFetch("GET", BaseURL+"/movies/"+id+"/extended?short=false", nil)
	if err != nil {
		return TvdbMovieResponse{}, err
	}

	var movieResp TvdbMovieResponse

	fetch.ExtractResponseJsonBody(resp, &movieResp)
	if err != nil {
		return TvdbMovieResponse{}, err
	}

	return movieResp, nil
}

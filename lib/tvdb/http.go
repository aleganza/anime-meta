package tvdb

import (
	"io"
	"net/http"
)

// Source - https://stackoverflow.com/a/54088988
// Posted by Jonathan Hall, modified by community. See post 'Timeline' for change history
// Retrieved 2026-07-28, License - CC BY-SA 4.0

func (c *Client) AuthenticatedHttpRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	return req, nil
}

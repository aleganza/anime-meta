package fetch

import (
	"encoding/json"
	"io"
	"net/http"
)

func Do(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Request(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func ExtractResponseJsonBody(resp *http.Response, dst any) error {
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

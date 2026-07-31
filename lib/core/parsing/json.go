package parsing

import (
	"encoding/json"
	"fmt"
	"io"
)

func DecodeJSON(body io.Reader, dst any) error {
	err := json.NewDecoder(body).Decode(dst)
	if err != nil {
		return fmt.Errorf("json decode error: %w", err)
	}

	return nil
}

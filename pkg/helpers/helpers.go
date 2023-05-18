package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

func BindRequestJSON(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
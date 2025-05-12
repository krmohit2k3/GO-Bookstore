package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// / What this function do?
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			//json.Unmarshal takes the body (converted to a []byte) and decodes it into x.
			return
		}
	}
}

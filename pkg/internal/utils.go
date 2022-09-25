package internal

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	t := "data"
	if code > 300 {
		t = "errors"
	}
	response, _ := json.Marshal(map[string]interface{}{
		t: payload,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (a *API) writeResponse(w http.ResponseWriter, body any, statusCode int) {
	enc, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(enc); err != nil {
		log.Println(fmt.Errorf("w.Write: %w", err))
	}
}

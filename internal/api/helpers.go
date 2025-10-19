package api

import (
	"emperror.dev/errors"
	"encoding/json"
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
		log.Println(errors.Wrap(err, "w.Write").Error())
	}
}

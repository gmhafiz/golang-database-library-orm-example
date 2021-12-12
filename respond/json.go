package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		http.Error(w, `{"message": "json error"}`, http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		http.Error(w, `{"message": "respond error"}`, http.StatusInternalServerError)
		return
	}
}

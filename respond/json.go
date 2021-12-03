package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)

	//(reflect.Ptr(reflect.ValueOf(payload)) && reflect.ValueOf(payload).Elem().IsNil())
	//if payload == nil ||
	//	(reflect.ValueOf(payload).Kind() == reflect.Ptr && reflect.ValueOf(payload).IsNil()) ||
	//	(reflect.ValueOf(payload).Elem().IsNil()) {
	//	return
	//}
	if payload == nil || &payload == nil{
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
		http.Error(w, `{"message": "wrote error"}`, http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"awesomeProject/urldata"
	"encoding/json"
	"net/http"
)

func GetData (data urldata.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := data.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}
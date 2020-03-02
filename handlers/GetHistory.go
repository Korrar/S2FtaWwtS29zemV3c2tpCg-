package handlers

import (
	"awesomeProject/historydata"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func GetHistory (data historydata.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		IdInt,err := strconv.Atoi(ID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		items := data.GetAll()
		for it :=range items{
			if items[it].Id == IdInt{
				items = append(items[:it], items[it+1:]...)
			}
		}
		json.NewEncoder(w).Encode(items)
	}
}
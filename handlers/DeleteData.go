package handlers

import (
	"awesomeProject/urldata"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func DeleteData (data urldata.Deleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		IdInt,err := strconv.Atoi(ID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		data.Delete(IdInt)
	}
}
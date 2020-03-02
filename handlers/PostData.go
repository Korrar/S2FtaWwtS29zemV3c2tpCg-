package handlers

import (
	"awesomeProject/historydata"
	"awesomeProject/urldata"
	"awesomeProject/workwork"
	"encoding/json"
	"net/http"
)
var i = 0
func PostData (data urldata.Adder ,hist historydata.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request := struct {
			Url  string `json:"URL"`
			Interval int `json:"Interval"`
		}{}
		err2 := json.NewDecoder(r.Body).Decode(&request)
		if err2 != nil  {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(413), 413)
			return
		}
		data.Add(urldata.Item{
			Id:  i,
			Url: request.Url,
			Interval: request.Interval,
		})
		i = i + 1
		go workwork.Working(i, request.Url , request.Interval , hist)


	}
}
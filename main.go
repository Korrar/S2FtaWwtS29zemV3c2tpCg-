package main

import (
	"awesomeProject/handlers"
	"awesomeProject/historydata"
	"awesomeProject/urldata"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)



func main() {
	port := ":8080"
	data := urldata.New()
	history := historydata.New()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Get("/api/fetcher" , handlers.GetData(data))
	r.Get("/api/fetcher/{ID}/history" , handlers.GetHistory(history))
	r.Post("/api/fetcher" , handlers.PostData(data , history))
	r.Delete("/api/fetcher/{ID}" , handlers.DeleteData(data))


	http.ListenAndServe(port, r)
}

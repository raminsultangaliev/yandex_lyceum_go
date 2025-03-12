package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Config struct {
	Host string
	Port string
}

type Router struct {
	config Config
	Router *mux.Router

	Handler Handler
}

func NewRouter(cfg Config, h *Handler) *Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", h.Get).Methods("GET")
	r.HandleFunc("/books", h.Add).Methods("POST")

	return &Router{config: cfg, Router: r}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", r.config.Host, r.config.Port), r.Router))
}

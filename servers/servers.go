package servers

import (
	"fmt"
	"net/http"
	"quote/data"
	"quote/handlers"

	"github.com/rs/cors"
)

type server struct {
	listenAddr string
	db         data.Database
	allowedIP  string
}

type Server interface {
	Start()
}

func NewServer(listenAddr string, db data.Database, allowedIP string) *server {
	return &server{
		listenAddr: listenAddr,
		db:         db,
		allowedIP:  allowedIP,
	}
}

func (s *server) Start() {

	router := http.NewServeMux()

	h := handlers.NewHandler(s.db)

	router.HandleFunc("POST /login", h.Login)
	router.HandleFunc("GET /welcome", h.Welcome)

	server := http.Server{
		Addr: s.listenAddr,
		Handler: cors.New(cors.Options{
			AllowedOrigins: []string{s.allowedIP},
		}).Handler(router),
	}

	fmt.Println("Start running server at port", server.Addr)
	server.ListenAndServe()
}

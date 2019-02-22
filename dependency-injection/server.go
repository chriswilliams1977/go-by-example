package main

import (
	"net/http"
)

//dependent on the PersonService and the Config
type Server struct {
	config        *Config
	personService *PersonService
}

//create a new handler
func (s *Server) Handler() http.Handler {
	//create empty mux
	mux := http.NewServeMux()

	//use the mux.Handle function to register this with our new ServeMux,
	//so it acts as the handler for all incoming requests with the URL path /people.
	mux.HandleFunc("/people", s.people)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *Config, service *PersonService) *Server {
	return &Server{
		config:        config,
		personService: service,
	}
}

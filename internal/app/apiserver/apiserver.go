package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start
func (s *APIServer) Start() error {
	if err := s.configueLogger(); err != nil {
		return err
	}

	s.configueRouter()

	s.logger.Info("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configueLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configueRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

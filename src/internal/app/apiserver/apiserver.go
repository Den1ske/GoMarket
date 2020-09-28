package apiserver

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/Den1ske/GoMarket/src/internal/app/store"
)

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter() 

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Start API server")

	return http.ListenAndServe(s.config.APIServerPort, s.router)
}

func (s *APIServer) configureLogger() error  {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/product", s.handleProduct())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st

	return nil
}

func (s *APIServer) handleProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `Hello`)
	}
}
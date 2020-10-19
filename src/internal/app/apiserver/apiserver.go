package apiserver

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/Den1ske/GoMarket/src/internal/app/store/sqlstore"
	"github.com/Den1ske/GoMarket/src/internal/app/controller"
)

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *sqlstore.Store
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
	s.router.HandleFunc("/products", controller.ListProducts())
}

func (s *APIServer) configureStore() error {

	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	s.store = sqlstore.New(db)

	return nil
}
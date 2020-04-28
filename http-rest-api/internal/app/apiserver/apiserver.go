package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start represents start of http server and connection to db
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting api server...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

// handleHello ... Идея возвращать в сигнатуре http.HandlerFunc дает возможность до return func(w http.ResponseWriter, r *http.Request)
//  мы можем определить переменные на 52-53, (type request struct ) которые будут использоваться только в этом хендлере. И этот код выполнится здесь один раз.
// Таким образом не захламляем код, а бизнес-логика будет внутри функции, которую мы возвращаем.
func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

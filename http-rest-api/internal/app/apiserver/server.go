package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lithium555/GolangStudying/http-rest-api/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  *sqlstore.Store
}

func newServer(store *sqlstore.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods(http.MethodPost)
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

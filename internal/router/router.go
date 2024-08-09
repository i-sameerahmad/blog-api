package router

import (
	"net/http"

	"github.com/ImranZahoor/blog-api/internal/controller"
	"github.com/gorilla/mux"
)

type (
	server struct {
		controller controller.Controller
		router     *mux.Router
	}
)

func NewServer(ctler controller.Controller) *server {
	return &server{
		router:     mux.NewRouter(),
		controller: ctler,
	}
}

func (s *server) RegisterHandlers() {
	// article handlers
	articleRouter := s.router.PathPrefix("/article").Subrouter()
	articleRouter.HandleFunc("/", s.controller.ListArticleHandler).Methods(http.MethodGet)
	articleRouter.HandleFunc("/{id}", s.controller.GetArticleByIDHandler).Methods(http.MethodGet)
	articleRouter.HandleFunc("/", s.controller.CreateArticleHandler).Methods(http.MethodPost)
	articleRouter.HandleFunc("/{id}", s.controller.UpdateArticleHandler).Methods(http.MethodPut)
	articleRouter.HandleFunc("/{id}", s.controller.DeleteArticleHandler).Methods(http.MethodDelete)
	//category handlers
	categoryRouter := s.router.PathPrefix("/category").Subrouter()
	categoryRouter.HandleFunc("/", s.controller.ListCategoryHandler).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/{id}", s.controller.GetCategoryByIDHandler).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/", s.controller.CreateCategoryHandler).Methods(http.MethodPost)
	categoryRouter.HandleFunc("/{id}", s.controller.UpdateCategoryHandler).Methods(http.MethodPut)
	categoryRouter.HandleFunc("/{id}", s.controller.DeleteCategoryHandler).Methods(http.MethodDelete)
	//user handlers
	userRouter := s.router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", s.controller.ListUserHandler).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", s.controller.GetUserByIDHandler).Methods(http.MethodGet)
	userRouter.HandleFunc("/", s.controller.CreateUserHandler).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", s.controller.UpdateUserHandler).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", s.controller.DeleteUserHandler).Methods(http.MethodDelete)

}

func (s *server) GetRouter() *mux.Router {
	return s.router
}

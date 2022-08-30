package controllers

import "github.com/selimaytac/TaskRegisterer/api/middlewares"

func (s *Server) InitializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddleWareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddleWareJSON(s.Login)).Methods("POST")

	// Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddleWareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{page}", middlewares.SetMiddleWareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddleWareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddleWareJSON(middlewares.SetMiddleWareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddleWareAuthentication(s.DeleteUser)).Methods("DELETE")
}

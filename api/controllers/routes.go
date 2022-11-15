package controllers

import "github.com/andreeabizu/xm-task/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Companies routes
	s.Router.HandleFunc("/company", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateCompany))).Methods("POST")
	s.Router.HandleFunc("/companies", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetCompanies))).Methods("GET")
	s.Router.HandleFunc("/company/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetCompany))).Methods("GET")
	s.Router.HandleFunc("/company/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCompany))).Methods("PUT")
	s.Router.HandleFunc("/company/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCompany)).Methods("DELETE")

}

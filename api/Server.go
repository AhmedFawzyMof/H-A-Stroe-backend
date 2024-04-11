package api

import (
	"HAstore/middleware"
	"HAstore/routes"
	"fmt"
	"net/http"
)

type Server struct {
	listenAddress string
}

func NewServer(listenAddress int) *Server {
	var Addr string = fmt.Sprintf(":%d", listenAddress)
	return &Server{
		listenAddress: Addr,
	}
}

func (s *Server) Start() error {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/home", routes.Home)
	mux.HandleFunc("GET /api/allproducts/{limit}", routes.AllProducts)
	mux.HandleFunc("GET /api/categories", routes.Categories)
	mux.HandleFunc("GET /api/category/{id}/{limit}", routes.CategoryPage)
	mux.HandleFunc("GET /api/subcategory/{id}/{limit}", routes.SubCategoryByid)
	mux.HandleFunc("GET /api/offer/{subcategory}/{limit}", routes.ProductByOffer)
	mux.HandleFunc("GET /api/product/{id}", routes.GetProduct)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	return http.ListenAndServe(s.listenAddress, middleware.CorsMiddleware(mux))
}

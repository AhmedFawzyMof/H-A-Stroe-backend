package api

import (
	"HAstore/router"
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

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (s *Server) Start() error {
	Router := router.NewRouter()
	Router.AddRoute("/navdata", GET, routes.NavData, false)
	Router.AddRoute("/allproducts", GET, routes.AllProducts, false)
	Router.AddRoute("/allcategories", GET, routes.AllCategories, false)
	Router.AddRoute("/filter", POST, routes.Filter, false)
	Router.AddRoute("/products/:slug", GET, routes.ProductBySlug, false)
	Router.AddRoute("/category/:name", GET, routes.ProductsByCategory, false)
	Router.AddRoute("/tags/:name", GET, routes.ProductsByTag, false)
	Router.AddRoute("/search/:query", GET, routes.ProductsBySearch, false)
	Router.AddRoute("/contact", POST, routes.ContactUs, false)
	Router.AddRoute("/register", POST, routes.Register, false)
	Router.AddRoute("/login", POST, routes.Login, false)
	Router.AddRoute("/checkout", POST, routes.CheckOut, false)
	Router.AddRoute("/auth/checkout", POST, routes.AuthCheckOut, true)
	Router.AddRoute("/order-histroy", GET, routes.OrderHistroy, true)

	http.HandleFunc("/", Router.Routes)
	return http.ListenAndServe(s.listenAddress, nil)
}

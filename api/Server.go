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

	Router.Insert("/navdata", routes.NavData, GET)
	Router.Insert("/allproducts", routes.AllProducts, GET)
	Router.Insert("/allcategories", routes.AllCategories, GET)
	Router.Insert("/filter", routes.Filter, POST)
	Router.Insert("/products/:slug", routes.ProductBySlug, GET)
	Router.Insert("/category/:id", routes.ProductsByCategory, GET)
	Router.Insert("/tags/:id", routes.ProductsByTag, GET)
	Router.Insert("/search/:query", routes.ProductsBySearch, GET)
	Router.Insert("/contact", routes.ContactUs, POST)
	Router.Insert("/register", routes.Register, POST)
	Router.Insert("/login", routes.Login, POST)
	Router.Insert("/checkout", routes.CheckOut, POST)
	Router.Insert("/auth/checkout", routes.AuthCheckOut, POST)
	Router.Insert("/order-histroy", routes.OrderHistroy, GET)
	Router.Insert("/offers", routes.Offers, GET)

	http.HandleFunc("/", Router.Router)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	return http.ListenAndServe(s.listenAddress, nil)
}

package router

import (
	"HAstore/middleware"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type routes []router

type handler func(res http.ResponseWriter, req *http.Request, params map[string]string)

type router struct {
	path    string
	method  string
	handler handler
	auth    bool
}

func NewRouter() *routes {
	return &routes{}
}

func (routes *routes) AddRoute(path, method string, handler handler, auth bool) {
	*routes = append(*routes, router{
		path:    path,
		method:  method,
		handler: handler,
		auth:    auth,
	})
}

func (routes *routes) Routes(res http.ResponseWriter, req *http.Request) {
	from := time.Now()
	for _, route := range *routes {
		if matched, params := match(route.path, req.URL.Path); matched {
			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", route.method)
			res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			if req.Method == "OPTIONS" {
				res.WriteHeader(http.StatusOK)
				return
			}
			if req.Method == route.method {
				if route.auth {
					Token := middleware.GetToken(req.Header.Get("Authorization"))
					email, err := middleware.VerifyToken(Token)
					if err != nil {
						middleware.SendError(err, res)
						return
					}

					params["authEmail"] = email
				}
				route.handler(res, req, params)
				since := time.Since(from)
				fmt.Println(since)
				return
			}

			since := time.Since(from)
			fmt.Println(since)
			http.Error(res, "Not Found", 404)
			return
		}
	}

	since := time.Since(from)
	fmt.Println(since)
	http.Error(res, "Not Found", 404)
}

func match(path, requestedPath string) (bool, map[string]string) {
	pathSlice := strings.Split(path, "/")
	requestedPathSlice := strings.Split(requestedPath, "/")
	params := make(map[string]string)

	if len(pathSlice) != len(requestedPathSlice) {
		return false, params
	}

	for i := 0; i < len(pathSlice); i++ {
		var pathName string = pathSlice[i]
		var requestedPathName string = requestedPathSlice[i]

		if strings.HasPrefix(pathName, ":") {
			paramsName := pathName[:1]
			params[paramsName] = requestedPathSlice[i]
		}

		if pathName != requestedPathName {
			return false, nil
		}

	}
	return true, params
}

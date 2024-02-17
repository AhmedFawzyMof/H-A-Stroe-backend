package router

import (
	"HAstore/middleware"
	"net/http"
	"strings"
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
	for _, route := range *routes {
		if matched, params := match(route.path, req.URL.Path); matched {
			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", route.method+", OPTIONS")
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
					route.handler(res, req, params)
					return
				}
				route.handler(res, req, params)
				return
			}

			http.Error(res, "Not Found", 404)
			return
		}
	}

	http.Error(res, "Not Found", 404)

}

func match(pattern, path string) (bool, map[string]string) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	if len(patternParts) == len(pathParts) {
		if pathParts[0] != patternParts[0] {
			return false, nil
		}
	}

	params := make(map[string]string)

	for i, patternPart := range patternParts {
		if strings.HasPrefix(patternPart, ":") {
			paramName := patternPart[1:]
			params[paramName] = pathParts[i]
		} else if patternPart != pathParts[i] {
			return false, nil
		}
	}

	return true, params
}

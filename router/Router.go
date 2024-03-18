package router

import (
	"HAstore/middleware"
	"fmt"
	"net/http"
	"time"
)

func (router *Trie) Router(res http.ResponseWriter, req *http.Request) {
	resultNode, isFound := router.Search(req.URL.Path)

	if !isFound {
		http.Error(res, "Not Found", 404)
		return
	}

	middleware.SetCors(res, resultNode.method)

	if req.Method == "OPTIONS" {
		res.WriteHeader(http.StatusOK)
		return
	}

	if req.Method == resultNode.method {
		now := time.Now()
		resultNode.handler(res, req, resultNode.params)
		since := time.Since(now)
		fmt.Println(since)
		return
	}
}

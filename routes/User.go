package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"io"
	"net/http"
)

func Register(res http.ResponseWriter, req *http.Request, params map[string]string) {
	if params["authEmail"] != "" {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var User models.Users

	if err := json.NewDecoder(req.Body).Decode(&User); err != nil {
		middleware.SendError(err, res)
		return
	}

	ResponseBytes := models.Users.Register(User, db)

	Response := make(map[string]interface{})

	if err := json.Unmarshal(ResponseBytes, &Response); err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func Login(res http.ResponseWriter, req *http.Request, params map[string]string) {
	if params["authEmail"] != "" {
		http.Error(res, "Forbidden", http.StatusForbidden)
		return
	}
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var User models.Users

	BodyData := req.Body
	Data, err := io.ReadAll(BodyData)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.Unmarshal(Data, &User); err != nil {
		middleware.SendError(err, res)
		return
	}

	ResponseBytes := models.Users.Login(User, db)

	Response := make(map[string]interface{})

	if err := json.Unmarshal(ResponseBytes, &Response); err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

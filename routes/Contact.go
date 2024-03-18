package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
)

func ContactUs(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var ContactUs models.Contact

	if err := json.NewDecoder(req.Body).Decode(&ContactUs); err != nil {
		middleware.SendError(err, res)
		return
	}

	ResponseBytes := models.Contact.AddContact(ContactUs, db)

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

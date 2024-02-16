package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"io"
	"net/http"
)

func ContactUs(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()

	defer db.Close()

	var ContactUs models.Contact

	BodyData := req.Body
	Data, err := io.ReadAll(BodyData)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.Unmarshal(Data, &ContactUs); err != nil {
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

package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
)

func Categories(res http.ResponseWriter, req *http.Request, params map[string]string) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	subcategory := models.SubCategory{}
	SubCategories, err := subcategory.GetAllSubCategory(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response:= make(map[string]interface{})
	Response["SubCategories"] = SubCategories

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
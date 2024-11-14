package utils

import (
	"encoding/json"
	"net/http"
	"tour_destination/model"
)

func Response(w http.ResponseWriter, statusCode int, message string, data any)  {
	
	response := model.Response {
		StatusCode: statusCode,
		Message: message,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func SuccessWithPage(w http.ResponseWriter, statusCode, page, totalData, limit, totalPage int, message string, data any)  {
	
	response := model.Response {
		StatusCode: statusCode,
		Page: page,
		TotalData: totalData,
		Limit: limit,
		TotalPage: totalPage,
		Message: message,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
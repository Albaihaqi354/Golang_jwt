package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponsWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponsWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "Success"

	if code >= 400 {
		status = "Failed"
	}

	if payload != nil {
		response = ResponsWithData{
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = ResponsWithoutData{
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}

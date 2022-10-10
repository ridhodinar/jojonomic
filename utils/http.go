package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error   bool        `json:"error"`
	ReffID  string      `json:"reff_id,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteSuccessResponse(w http.ResponseWriter, reffID string) {
	resp := response{
		ReffID: reffID,
	}

	data, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func WriteSuccessResponseWithData(w http.ResponseWriter, data interface{}, reffID string) {
	resp := response{
		ReffID: reffID,
		Data:   data,
	}

	output, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func WriteErrorResponse(w http.ResponseWriter, reffID string, err error) {
	resp := response{
		Error:   true,
		ReffID:  reffID,
		Message: err.Error(),
	}

	output, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

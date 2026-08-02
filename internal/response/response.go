package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success    bool               `json:"success"`
	Message    string             `json:"message,omitempty"`
	Data       any                `json:"data,omitempty"`
	Pagination PaginationResponse `json:"pagination,omitzero"`
	Error      string             `json:"error,omitempty"`
}

type PaginationResponse struct {
	CurrentPage  int  `json:"current_page"`
	PerPage      int  `json:"per_page"`
	TotalResults int  `json:"total_results"`
	TotalPages   int  `json:"total_pages"`
	HasPrev      bool `json:"has_prev"`
	HasNext      bool `json:"has_next"`
}

func Error(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error:   err,
	})
}

func Json(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func PaginatedJson(w http.ResponseWriter, status int, message string, data any, pagination PaginationResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

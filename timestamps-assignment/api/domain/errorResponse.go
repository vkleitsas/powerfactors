package domain

type ErrorResponse struct {
	Status      string `json:"status"`
	Description string `json:"desc"`
}

package models

type ListImagesByBreedResponse struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

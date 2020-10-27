package models

type ListAllBreedsResponse struct {
	Message map[string][]string `json:"message"`
	Status  string              `json:"status"`
}

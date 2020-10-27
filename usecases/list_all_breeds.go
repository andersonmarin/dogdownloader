package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andersonmarin/dogdownloader/models"
)

func ListAllBreeds() (map[string][]string, error) {
	url := "https://dog.ceo/api/breeds/list/all"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response models.ListAllBreedsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("error %s", response.Status)
	}

	return response.Message, nil
}

package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andersonmarin/dogdownloader/models"
)

func ListImagesByBreed(breed string) ([]string, error) {
	url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images", breed)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response models.ListImagesByBreedResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("error %s", response.Status)
	}

	return response.Message, nil
}

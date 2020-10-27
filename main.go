package main

import (
	"sync"

	"github.com/andersonmarin/dogdownloader/usecases"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(6)

	breeds := make(chan string)
	images := make(chan string)

	go listAllBreedsQueue(&wg, breeds)
	go listImagesListener(&wg, breeds, images)
	go downloadImageListener(&wg, images)
	go downloadImageListener(&wg, images)
	go downloadImageListener(&wg, images)
	go downloadImageListener(&wg, images)

	wg.Wait()
}

func listAllBreedsQueue(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()

	breeds, err := usecases.ListAllBreeds()
	if err != nil {
		panic(err)
	}
	for breed := range breeds {
		c <- breed
	}
	close(c)
}

func listImagesListener(wg *sync.WaitGroup, breeds, imagesChan chan string) {
	defer wg.Done()
	for breed := range breeds {
		images, err := usecases.ListImagesByBreed(breed)
		if err != nil {
			panic(err)
		}

		for _, image := range images {
			imagesChan <- image
			break
		}
	}
	close(imagesChan)
}

func downloadImageListener(wg *sync.WaitGroup, images chan string) {
	defer wg.Done()

	for image := range images {
		err := usecases.DownloadImage(image)
		if err != nil {
			panic(err)
		}
	}
}

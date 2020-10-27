package usecases

import "sync"

func ListImagesListener(wg *sync.WaitGroup, breeds, imagesChan chan string) {
	defer wg.Done()
	for breed := range breeds {
		images, err := ListImagesByBreed(breed)
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

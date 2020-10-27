package usecases

import "sync"

func DownloadImageListener(wg *sync.WaitGroup, images chan string) {
	defer wg.Done()

	for image := range images {
		err := DownloadImage(image)
		if err != nil {
			panic(err)
		}
	}
}

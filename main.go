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

	go usecases.ListAllBreedsQueue(&wg, breeds)
	go usecases.ListImagesListener(&wg, breeds, images)
	go usecases.DownloadImageListener(&wg, images)
	go usecases.DownloadImageListener(&wg, images)
	go usecases.DownloadImageListener(&wg, images)
	go usecases.DownloadImageListener(&wg, images)

	wg.Wait()
}

package usecases

import (
	"fmt"
	"sync"
)

func ListImagesListener(wg *sync.WaitGroup, breeds, c chan string) {
	defer wg.Done()
	for breed := range breeds {
		fmt.Printf("fetching images list for %s...", breed)

		images, err := ListImagesByBreed(breed)
		if err != nil {
			panic(err)
		}

		fmt.Println("done")

		for i, image := range images {
			c <- image
			fmt.Printf("\rdownloading images for %s... %.2f%%", breed, (float64(i+1)/float64(len(images)))*100)
		}

		fmt.Println()
	}
	close(c)
}

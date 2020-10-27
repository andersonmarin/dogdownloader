package usecases

import (
	"fmt"
	"sync"
)

func ListAllBreedsQueue(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()

	fmt.Print("fetching breeds list...")

	breeds, err := ListAllBreeds()
	if err != nil {
		panic(err)
	}

	fmt.Println("done")

	for breed := range breeds {
		c <- breed
	}
	close(c)
}

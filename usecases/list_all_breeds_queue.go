package usecases

import "sync"

func ListAllBreedsQueue(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()

	breeds, err := ListAllBreeds()
	if err != nil {
		panic(err)
	}
	for breed := range breeds {
		c <- breed
	}
	close(c)
}

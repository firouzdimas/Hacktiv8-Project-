package main

import (
	"fmt"
	"sync"
)

func main() {
	coba := []string{"coba1", "coba2", "coba3"}
	bisa := []string{"bisa1", "bisa2", "bisa3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go pritnHasil(i, coba, &wg)
		wg.Add(1)
		go pritnHasil(i, bisa, &wg)
	}

	wg.Wait()
}

func pritnHasil(index int, slice []string, wg *sync.WaitGroup) {
	fmt.Println(slice, index)
	wg.Done()
}
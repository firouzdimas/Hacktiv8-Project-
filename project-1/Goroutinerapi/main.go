package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	coba := []interface{}{"coba1", "coba2", "coba3"}
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go printHasil(i, coba, bisa, &wg, &mtx)
	}

	wg.Wait()
}

func printHasil(index int, coba []interface{}, bisa []interface{}, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	fmt.Println(coba, index)
	fmt.Println(bisa, index)
	mtx.Unlock()
	wg.Done()
}

package main

import "fmt"

func main() {
	var kata string = "selamat malam"
	
	for _, letter := range kata {
		fmt.Printf("%c \n", letter)
	}

	hitungkata := make(map[string]int)
	for _, char := range kata {
	hitungkata[string(char)]++
	}

	fmt.Println(hitungkata)
	
}

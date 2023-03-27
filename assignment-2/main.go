package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j < 5; j++ {
		fmt.Println("Nilai j = ", j)
	}
	String := "САШАРВО"
	for i, v := range String {
		fmt.Printf("character %#U  start at byte position %v\n", v, i)
	}
	j := 5
	for{j++
		if j <= 10 {
			fmt.Println("Nilai j = ", j)
		} else {
			break
		}
	}
}
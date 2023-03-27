package main

import ("fmt"
)

func main() {
	//menampilkan nilai i : 21 
	i := 21
   	fmt.Printf("%d\n", i)

	//menampilkan tipe data dari variabel i
    fmt.Printf("%T\n", i)

	//menampilkan tanda %
    fmt.Println("%")

	//menampilkan nilai boolean j : true
    j := true
    fmt.Printf("%t\n\n", j)

	//menampilkan nilai base 2 dari i : 10101
    fmt.Printf("%b \n", i)

	//menampilkan unicode Rusia Я
	fmt.Printf("%U \n\n", 'Я')
    
	// menampilkan nilai base 10 dari i : 21
    fmt.Printf("%d \n", i)       

	// menampilkan nilai base 8 dari i : 25
	fmt.Printf("%o \n", i)      

	// menampilkan nilai base 16 : f
	fmt.Printf("%x \n", 15)   

	// menampilkan nilai base 16 : F
	fmt.Printf("%X \n", 15) 
	
	// menampilkan unicode karakter Я : U+042F
	fmt.Printf("%c \n", '\u042F') 

	// menampilkan nilai variabel k float64 = 123.456
	var k float64 = 123.456

	// menampilkan float : 123.456000
	fmt.Printf("%f \n", k)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%E \n", k)
}
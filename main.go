package main

import "fmt"

func main() {
	// // fmt.Println("Hello, world")

	// // fmt.Println("2 x 1 =", 2*1)
	// // fmt.Println("2 x 2 =", 2*2)
	// // fmt.Println("2 x 3 =", 2*3)
	// // // 2 x 1 = 2 

	// // x := 1
	// // y := 2
	// // fmt.Println("x + y =", x+y)

	// for i := 1; i<=10; i++ {
	// 	fmt.Println(i)
	// }

	// a := 1
	// fmt.Println(a)

	for i:= 2; i<= 9; i++ {
		jaeuk(i)
	}

	
	

}

func jaeuk(a int) {
	for i:=1; i<=9; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}

// Make a calculator
func calculator(a int, b int) {
	fmt.Println(a, "+", b, "=", a+b)
	fmt.Println(a, "-", b, "=", a-b)
	fmt.Println(a, "*", b, "=", a*b)
	fmt.Println(a, "/", b, "=", a/b)
}



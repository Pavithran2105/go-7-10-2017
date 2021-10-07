package main

import "fmt"

func action(a int, b int) {
	for k := a; k <= b; k++ {
		if k%2 == 0 {
			fmt.Println(k)
		}
	}

}

func main() {
	var num1 int
	fmt.Println("This program will give even number for the selected numbers")
	fmt.Println("enter the the start number")

	fmt.Scan(&num1)
	var num2 int
	fmt.Println("enter the the end number")
	fmt.Scan(&num2)
	action(num1, num2)

	// for k := num1; k <= num2; k++ {
	// 	if k%2 == 0 {
	// 		fmt.Println(k)
	// 	}
	// }
}

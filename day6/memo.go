package main

import "fmt"

func swap(a, b *int) {
	// we get the address not values so using (*)operator
	// & by using * operator we can change or modify the values reside in the memory.
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 5, 10
	fmt.Println("Before swap: ",x,y)
	swap(&x, &y) //passing the memory address
	fmt.Println("After swap: ",x,y)
}

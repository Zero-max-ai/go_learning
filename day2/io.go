package main

import "fmt"

func fuck_ios() {
	var (
		line string
		age int = 21
		height float32 = 5.11
	)
	
	fmt.Print("enter a line: ")
	fmt.Scanln(&line)
	fmt.Printf("You entered: %s\n", line)

	formatted := fmt.Sprintf("Age: %d, Height: %.2f", age, height)
	fmt.Printf(formatted)
}

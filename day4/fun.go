package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiple_add(a,b int, c,d float32) (int, float32) {
	return a + b, c + d
}

func main() {
	fmt.Printf("6 + 7: %d\n",add(6,7))
	intRes, floatRes := multiple_add(1,6,2.03,6.23)
	fmt.Printf("1 + 6: %d, 2.03 + 6.23: %.2f\n", intRes, floatRes)
	
	res1, err := divide(5.2, 0);
	if (err != nil) {
		fmt.Println("Error: ", err)
	}
	fmt.Println("%.2f", res1)

	res, err1 := divide(4.1, 2.3)
	if (err1 != nil) {
		fmt.Println("Error: ", err1)
	}
	fmt.Println("Result of 4.1 / 2.3: %.2f", res)

	content, err2 := readFile("test.txt")
	if err2 != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("File content: %s\n", content)
	}
}


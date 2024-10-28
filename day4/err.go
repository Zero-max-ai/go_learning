package main

import "fmt"

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("division error: cannot divide by zero")
	}
	return a / b, nil
}

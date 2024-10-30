package main

import (
	"fmt"
	"os"
	"log"
	"errors"
)

func openFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buffer := make([]byte, 100)
	bytesRead, err := f.Read(buffer)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("File contents: ", string(buffer[:bytesRead]))

}

func safeDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from Panic: %v\n", r)
		}
	}()

	fmt.Println(a / b)
}

func divide(a, b float32) (result float32, err error) {
	if b == 0 {
		err = errors.New("cannot divide by Zero")
		return
	}

	result = a/b
	return
}

func main() {
	// defer function
	// openFile("test.txt")
	// openFile("test1.txt")

	// panic & recover
	// safeDivision(4, 2)
	// safeDivision(4, 0)
	
	res, err := divide(10,2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("10/2 = ", res)
	}

	res, err = divide(10,0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("10/0 = ", res)
	}
}

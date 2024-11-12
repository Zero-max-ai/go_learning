package main

import (
	"fmt"
	"errors"
)

// 1. a function which does what i want
func sayHello() {
    fmt.Println("Hello!")
}

// 2. a function with parameter
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// 3. a function with return value (single)
func add(x, y int) int {
    return x + y
}

// 4. a function with Multiple Return Values
func divide(x, y float32) (float32, error) {
    if y == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return x / y, nil
}

// 5. a function with Named Return Values
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return because we used named return values
}

// 6. a function with deferred Execution
func defer_exe() {
	defer fmt.Println("No, I came first :)")
    fmt.Println("I came first ;)")
}

// 7.) a function which is anonymous
func closure() func () int {
	count := 0
	return func() int {
		count ++
		return count
	}
}

func main() {
	
	var name string = "akshat"
	var age int = 21

	var (
		location = "India, IN"
		hobbies []string = []string{"IT", "Astrology", "Gaming", "Basketball"}
	)

	const VERSION = "1.1"

	fmt.Printf("My name is: [%s], I am [%d] years old. I live in [%s]\n", name , age, location)
	fmt.Print("My hobbies are: -")
	for _, work := range hobbies {
		fmt.Printf("%s -", work)
	}
	
	// if statements 
	
	if age >= 18 {
		fmt.Println("I am an adult & do have driving license.")
	} else {
		fmt.Println("I am a minor ;)")
	}

	// switch case
	gender := "male"

	switch gender {
		case "male": fmt.Println("I am male")
		break
		case "female": fmt.Println("I am female")
			break
		default: fmt.Println("I know its not default, but why are you gay??")
		break
	}

	// for loops  have multiple types, we can also use it like while loop and iterative as well
	for i:=0; i<5; i++ {
		fmt.Println("Iterative: ", i)
	}
	i := 0
	for {
		if i == 5 { 
			break
		}
		fmt.Println("While type for loop: ", i)
		i ++;
	}

	array_type := []int{7,8,6,5,43,2}
	for index, val := range array_type {
		fmt.Printf("This value [%d] is at this index: [%d]\n", val, index)
	}

	// functons calling
	sayHello()

	greet("akshat")

	fmt.Println("The result in add is: 4 + 5: ", add(4,5))
	
	quo, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("quo is: ", quo)
	}

	area, perim := rectangle(3.0, 4.0)
	fmt.Println("area is 	%.2f and perimeter is: %.2f", area, perim)

	defer_exe()

	counter := closure()
	fmt.Println(counter())
	fmt.Println(counter())

}

package main

import (
	"fmt"
	"strings"
)

func main() {

	// arrays
	// init empty which by default have 0s init
	var numbers [5]int
	fmt.Printf("Empty array: %v\n", numbers)

	// init with values
	scores := [5]int{0,8,5,3,2}
	fmt.Printf("Scores: %v\n", scores)

	days := [...]string{"Monday", "Tuesday", "Wednesday"}
	fmt.Printf("Days: %v, Length %d\n", days, len(days))

	scores[1] = 89
	fmt.Println("Updated score: %v\n", scores[1])

	// slices
	numbers2 := make([]int, 5, 10)
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", numbers2, len(numbers2), cap(numbers2))

	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruits: %v\n", fruits)

	fruits = append(fruits, "grape")
	fmt.Printf("After append: %v\n", fruits)

	subset := fruits[1:3]
	fmt.Printf("Subset: %v\n", subset)

	// maps
	// map init without any values and then adding 1 by 1
	studentScores := make(map[string]int)
	studentScores["akshat"] = 95
	studentScores["piyuss"] = 87
	studentScores["guts"] = 85
	fmt.Printf("Student scores: %v\n", studentScores)

	// map init with values
	ages := map[string]int {
		"akshat": 21,
		"supriya": 20,
		"sakshi": 21,
	}

	fmt.Printf("Ages init: %v\n", ages)

	age, exists := ages["akshat"]
	if exists {
		fmt.Printf("akshat's age: %d\n", age)
	}

	// deleting entry from map
	delete(ages, "supriya")
	fmt.Printf("After deletion: %v\n", ages)


	// strings and runes
	str := "Hello"
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length in bytes: %d\n", len(str))

	// coverting string to runes
	runes := []rune(str)
	fmt.Printf("Length in runes: %d\n", len(runes))

	fmt.Print("Bytes: ")
	for i:=0; i<len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()

	fmt.Print("Runes: ")
	for _, r := range str {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	text := "  Hello, World  "
	fmt.Printf("Original Text: %q\n", text)
	fmt.Printf("Trimmed Text: %q\n", strings.TrimSpace(text))
	fmt.Printf("Upper Case: %s\n", strings.ToUpper(text))
	fmt.Printf("Check if Text Contains 'World': %v\n", strings.Contains(text, "World"))
}

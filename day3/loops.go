package main

import "fmt"

func forloops() {
for i:=0; i<=5; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 10 {
		fmt.Printf("%d x %d = %d\n",2,i,2*i)
		i++;
	}
}
func ifstat() {
	gender := 'M' // can't change na :)
	if gender == 'M' {
		fmt.Println("its a male")
	} else if gender == 'F' {
		fmt.Println("its a female")
	} else {
		fmt.Println("its transformer")
	}
}
func switchs() {
	num := 2
	switch num {
		case 0 : fmt.Println("Its 0");
		case 1 : fmt.Println("Its 1")
		default : fmt.Println("Neither 0 nor 1")
	}
}

func main() {

	forloops()
	ifstat()
	switchs();

}

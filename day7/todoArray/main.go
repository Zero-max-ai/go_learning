package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// create a Todo type structure
type Todo struct {
	ID					int
	Task				string
	Completed		bool
}

var todos[] Todo
var nextId int = 0

// a function to add tasks in an array
func addTodo(taskName string) {
	todos = append(todos, Todo{ID: nextId, Task: taskName, Completed: false})
	nextId++
}

func listTasks() {
	if len(todos) == 0 {
		fmt.Println("task queue is empty!!")
		return
	}
	for _, todo := range todos {
		status := " "
		if (todo.Completed) {
			status = "c"
		}
		fmt.Printf("[%s] %d: %s", status, todo.ID, todo.Task)
	}
	fmt.Println("")
}

func changeTaskStatus(id int) {
	for index, todo := range todos {
		if (todo.ID == id) {
			todos[index].Completed = !todo.Completed
			return
		}
	}
	fmt.Println("id not found: ", id)
}

func deleteTask(id int) {
	for index, todo := range todos {
		if (todo.ID == id) {
			todos = append(todos[:index], todos[index+1:]...)
			fmt.Printf("task: %s is deleted with id: %d\n", todo.Task, todo.ID)
			return
		}
	}
	fmt.Println("id not found: ", id)
}

func main() {
	// creating reader instance of bufio
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter (add, list, complete, delete, exit): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		switch command {
			case "add":
				fmt.Print("Enter the task: ")
				task, _ := reader.ReadString('\n')
				task = strings.TrimSpace(task)
				addTodo(task)
			case "list":
				listTasks()
			case "complete":
				fmt.Print("Enter the id: ")
				id := -1
				fmt.Scanf("%d", &id)
				changeTaskStatus(id)
			case "delete":
				fmt.Print("Enter the id: ")
				id := -1
				fmt.Scanf("%d", &id)
				deleteTask(id)
			case "exit":
				fmt.Println("exiting the program....")
				return
			default:
				fmt.Println("unknown command!! please provide valid input...")
		}
	}
}

package main

import "fmt"

func main() {

	fmt.Println("Task Tracker CLI")

	var command string
	var tasks []string

	for {

		fmt.Print("Enter Command: ")
		fmt.Scanln(&command)
		var newTask string

		switch command {

		case "add":
			fmt.Print("Enter task: ")
			fmt.Scanln(&newTask)
			fmt.Print("Adding task...\n")
			tasks = append(tasks, newTask)

			fmt.Print("Task Added! \n")

		case "remove":
			var taskNum int
			fmt.Print("Enter task number to delete: ")
			fmt.Scanln(&taskNum)
			index := taskNum - 1
			if index < 0 || index >= len(tasks) {
				fmt.Println("Invalid task number")
			} else {
				tasks = append(tasks[:index], tasks[index+1:]...)
			}
			fmt.Print("Removing task...\n")

		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
			} else {
				fmt.Print("Listing tasks...\n")
				for index, task := range tasks {
					fmt.Println(index+1, task)
				}
			}

		case "exit":
			fmt.Print("exiting...")
			return

		default:
			fmt.Print("Invalid command, try again\n")

		}

	}

}

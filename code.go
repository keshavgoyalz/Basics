package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID   int
	NAME string
}

func main() {
	var tasks []Task
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Welcome to the Task Manager")
		fmt.Println("1 Add task")
		fmt.Println("2 Update task")
		fmt.Println("3 View All tasks")
		fmt.Println("4 View Task by ID or NAME")
		fmt.Println("5 Delete task by ID or NAME")
		fmt.Println("6 Exit")

		fmt.Println("Enter the number:")
		input.Scan()
		selected := input.Text()

		if selected == "1" {
			fmt.Print("Task Name: ")
			input.Scan()
			name := input.Text()
			id := len(tasks) + 1
			task := Task{
				ID:   id,
				NAME: name,
			}
			tasks = append(tasks, task)
			fmt.Println("New Task Added")
		}

		if selected == "2" {
			fmt.Print("Enter Task ID: ")
			input.Scan()
			id := input.Text()
			taskID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid task ID!")
				fmt.Println()
				continue
			}

			itemFound := false
			for i := range tasks {
				if tasks[i].ID == taskID {
					fmt.Printf("Enter new name for task %d: ", taskID)
					input.Scan()
					name := input.Text()
					tasks[i].NAME = name
					itemFound = true
					break
				}
			}

			if itemFound {
				fmt.Println("Task updated successfully!")
			} else {
				fmt.Println("Task not found!")
			}
			fmt.Println()

		}

		if selected == "3" {
			fmt.Println("Tasks:")
			for _, task := range tasks {
				fmt.Printf("%d = %s\n", task.ID, task.NAME)
			}
			fmt.Println()
		}

		if selected == "4" {
			fmt.Print("Enter task ID or task Name to view: ")
			input.Scan()
			input := input.Text()

			found := false
			for _, task := range tasks {
				if strconv.Itoa(task.ID) == input || task.NAME == input {
					fmt.Printf("Task ID: %d\n", task.ID)
					fmt.Printf("Task Name: %s\n", task.NAME)
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Task not found!")
			}
			fmt.Println()
		}

		if selected == "5" {
			fmt.Print("Enter task ID or task Name to delete: ")
			input.Scan()
			input := input.Text()
			found := false
			for i, task := range tasks {
				if strconv.Itoa(task.ID) == input || task.NAME == input {
					tasks = append(tasks[:i], tasks[i+1:]...)
					found = true
					break
				}
			}
			if found {
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Task not found!")
			}
			fmt.Println()
		}

		if selected == "6" {
			os.Exit(0)
		}
		fmt.Print("Do you want to continue? (y/n): ")
		input.Scan()
		continueChoice := input.Text()

		if strings.ToLower(continueChoice) != "y" {
			break
		}

	}
}

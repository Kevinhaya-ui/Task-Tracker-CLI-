package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: please provide a command!")
		return
	}

	command := os.Args[1]
	fmt.Printf("Success! You entered a command: %s\n", command)

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: please provide a configuration to add!")
			return
		}

		jsonData, err := os.ReadFile("task.json")
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}
		fmt.Println("Succesfully read the task")

		var task []Task
		err = json.Unmarshal(jsonData, &task)
		if err != nil {
			fmt.Println("Error decoding JSON: ", err)
			return
		}
		fmt.Printf("Succesfull, file parsed and found %d existing tasks\n", len(task))

		newTask := Task{
			ID:          len(task) + 1,
			Description: os.Args[2],
			Status:      "pushed",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		task = append(task, newTask)

		jsonData, err = json.MarshalIndent(task, "", "  ")
		if err != nil {
			fmt.Println("Error detected", err)
		}

		err = os.WriteFile("task.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error occured while writing file")
			return
		}
		fmt.Println("Succesfully updated the file on Hard Drive")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: please provide a configuration to delete!")
			return
		}

		jsonData, err := os.ReadFile("task.json")
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}
		fmt.Println("Succesfully read the task")

		var task []Task
		err = json.Unmarshal(jsonData, &task)
		if err != nil {
			fmt.Println("Error decoding JSON: ", err)
			return
		}
		fmt.Printf("Succesfull, file parsed and found %d existing tasks\n", len(task))

		newTask := Task{
			ID:          len(task) + 1,
			Description: os.Args[2],
			Status:      "deleted",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		task = append(task, newTask)

		jsonData, err = json.MarshalIndent(task, "", "  ")
		if err != nil {
			fmt.Println("Error detected", err)
		}

		err = os.WriteFile("task.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error occured: ", err)
			return
		}
		fmt.Println("Succesfully wrote to Hard Drive")

	case "buy":
		if len(os.Args) < 3 {
			fmt.Println("Error: please provide an item to buy!")
			return
		}

		jsonData, err := os.ReadFile("task.json")
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}
		fmt.Println("succesfully read the task")

		var task []Task
		err = json.Unmarshal(jsonData, &task)
		if err != nil {
			fmt.Println("Error decoding JSON", err)
			return
		}
		fmt.Printf("Succesfull, file parsed and found %d existing tasks\n", len(task))

		newTask := Task{
			ID:          len(task) + 1,
			Description: os.Args[2],
			Status:      "Order Confimed",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		task = append(task, newTask)
		jsonData, err = json.MarshalIndent(task, "", "  ")
		if err != nil {
			fmt.Println("Error occured", err)
		}
		err = os.WriteFile("task.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}
		fmt.Println("succesfully wrote file to Hard Drive")

	default:
		fmt.Println("Unknown command.")
	}
}

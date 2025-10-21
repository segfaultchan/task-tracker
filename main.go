package main

import (
	//	"encoding/json"
	"fmt"
	"strings"
	"time"
	//	"errors"
	"bufio"
	"os"
)

const (
	PROMPT = "tasks-$ "
	VERSION = "0.0.1"
)

func main() {
	tasks := make([]Task, 0)
	run := true
	for run{
		input,_ := smartInput(PROMPT)

		switch input[0] {
			case "exit":
				run = false
			case "help":
				helpT()
			case "list":
				listT(tasks)
			case "create":
				createT(&tasks, input[1],input[2])
			default:
				fmt.Println("unknown command")
		}
	}
}

type Task struct {
	desc string
	status string
	createdAt time.Time
	updatedAt time.Time
}

func createT (tasks *[]Task, desc string, status string) (error) {
	var task Task
	task.desc = desc
	task.status = status
	task.createdAt = time.Now()
	task.updatedAt = task.createdAt
	*tasks = append(*tasks, task)
	return nil
}

func listT (tasks []Task) (error) {
	for i,t := range tasks {
		fmt.Println("task id:", i)
		fmt.Println("{")
		fmt.Println("\tdescription:\t", t.desc )
		fmt.Println("\tstatus:\t\t", t.status )
		fmt.Println("\tcreated at:\t\t", t.createdAt )
		fmt.Println("\tupdated at:\t\t", t.updatedAt )
		fmt.Println("}")
	}
	return nil
}

func greeting() {
	fmt.Printf("Task Tracker v%v\n",VERSION)
}

func helpT() {
	fmt.Println("help:")
	fmt.Println("\t'list' - list of tasks")
	fmt.Println("\t'create.description.status' - create task")
	fmt.Println("\t'exit' - exit program")
	fmt.Println()
}

// function fot safety input and with prompt for tasks
func smartInput(prompt ...string) ([]string, error) {
	if len(prompt) == 0 {
		fmt.Print("input-$ ")
	} else {
	fmt.Print(prompt[0])
	}
	// output for return tokens
	var output = make([]string,3)
	// input a command
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	output = strings.Split(input, ".")

	return output,nil
}


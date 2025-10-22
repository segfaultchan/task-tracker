package main

import (
	//	"encoding/json"
	"fmt"
	"strings"
	"time"
	//	"errors"
	"bufio"
	"os"
	"strconv"
)

const (
	PROMPT = "tasks-$ "
	VERSION = "0.0.1"
)

func main() {
	greeting()
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
			case "delete":
				arg_int,err := strconv.Atoi(input[1])
				if (err != nil) {
					fmt.Println("error, not int argument")
					continue
				}
				deleteT(&tasks, arg_int)
			case "update":
				arg_int,err := strconv.Atoi(input[1])
				if (err != nil) {
					fmt.Println("error, not int argument")
					continue
				}
				updateT(tasks,arg_int,input[2])
				fmt.Println("status updated")
			case "print":
				arg_int,err := strconv.Atoi(input[1])
				if (err != nil) {
					fmt.Println("error, not int argument")
					continue
				}
				printT(tasks,arg_int)
				
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
	fmt.Println("task has created")
	return nil
}

func listT (tasks []Task) (error) {
	if len(tasks) == 0 {
		fmt.Println("no tasks there")
		return nil
	}
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

func deleteT (tasks *[]Task, index int) error {
	if len(*tasks) < index+1 {
		fmt.Println("not existing index of slice")
		return nil
	}
	same := *tasks
	*tasks = append(same[:index], same[index+1:]...)
	fmt.Printf("%v element has deleted\n", index)
	return nil
}

func updateT (tasks []Task, index int, status string) error {
	if len(tasks) < index+1 {
		fmt.Println("not existing index of slice")
		return nil
	}
	ptr := &tasks[index]
	ptr.status = status
	ptr.updatedAt = time.Now()
	return nil
}

func printT (tasks []Task, index int) error {
	if len(tasks) < index+1 {
		fmt.Println("not existing index of slice")
		return nil
	}
	t := tasks[index]
	fmt.Println("task id:", index)
	fmt.Println("{")
	fmt.Println("\tdescription:\t", t.desc )
	fmt.Println("\tstatus:\t\t", t.status )
	fmt.Println("\tcreated at:\t\t", t.createdAt )
	fmt.Println("\tupdated at:\t\t", t.updatedAt )
	fmt.Println("}")

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
	output = strings.Split(input, " ")

	return output,nil
}

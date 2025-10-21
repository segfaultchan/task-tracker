package main

import (
//	"encoding/json"
	"fmt"
	"time"
	"errors"
)

const (
	PROMPT = "tasks-$ "
	VERSION = "0.0.1"
)

func main() {
	input,_ := smartInput(PROMPT)
	for _,v := range input {
		fmt.Println("value: ", v)
	}
/*	tasks := make([]Task, 0)
	p1,p2,p3,_ := smartInput(PROMPT)
	run := true
	for run{
		switch p1 {
			case "exit":
				run = false
			case "list":
				listT(tasks)
			case "create":
				createT(tasks, p2,p3)
			default:
				fmt.Println("unknown command")
		}
	}*/
}

type Task struct {
	desc string
	status string
	createdAt time.Time
	updatedAt time.Time
}

func createT (tasks []Task, desc string, status string) (error) {
	var task Task
	task.desc = desc
	task.status = status
	task.createdAt = time.Now()
	task.updatedAt = task.createdAt
	tasks = append(tasks, task)
	return nil
}

func listT (tasks []Task) (error) {
	for i,t := range tasks {
		fmt.Println("task id:", i)
		fmt.Println("{")
		fmt.Println("\tdescription:\t", t.desc )
		fmt.Println("\tstatus:\t\t", t.status )
		fmt.Println("\tcreated at:\t\t", t.status )
		fmt.Println("\tupdated at:\t\t", t.status )
		fmt.Println("}")
	}
	return nil
}

func greeting() {
	fmt.Printf("Task Tracker v%v\n",VERSION)
}

// function fot safety input and with prompt for tasks
func smartInput(prompt ...string) ([]string, error) {
	if len(prompt) == 0 {
		fmt.Print("input-$ ")
	} else {
	fmt.Print(prompt[0])
	}
	var output = make([]string,3)
	_, e := fmt.Scan(&output[0],&output[1],&output[2])
	if e != nil {
		e = errors.New("input error")
		return nil,e
	}
	return output,nil
}


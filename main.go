package main

import (
	"fmt"
	"time"
	"errors"
)

const (
	PROMPT = "tasks-$ "
	VERSION = "0.0.1"
)

type Task struct {
	id int
	desc string
	isDone bool
	createdAt time.Time
	updatedAt time.Time
}

type Commands int
const (
	Help Commands = iota
	List
	Create
	Delete
	Edit
	Error
)

func greeting() {
	fmt.Printf("Task Tracker v%v\n",VERSION)
}

// function fot safety input and with prompt for tasks
func smartInput(prompt ...string) (string, error) {
	if len(prompt) == 0 {
		fmt.Print("input-$ ")
	} else {
	fmt.Print(prompt[0])
	}
	var input string
	_, e := fmt.Scanln(&input)
	if e != nil {
		e = errors.New("input error")
		return "", e
	}
	return input, nil
}

/*func parseInput(input string) (Commands, error) {
	switch input {
	case "help":
		return Help, nil
	case "list":
		return List, nil
	case "create":
		return Create, nil
	case "delete":
		return Delete, nil
	case "edit":
		return Edit, nil
	default:
		return Error, errors.New("unknown command")
	}
}*/

func main() {
	greeting()
//	input, _ := smartInput(PROMPT)
//	parseInput(input)
}

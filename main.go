package main

import (
	"fmt"
	"time"
	"errors"
)

const (
	PROMPT = "tasks-$ "
)

type Task struct {
	id int
	desc string
	isDone bool
	createdAt time.Time
	updatedAt time.Time
}

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

func main() {
	input, _ := smartInput(PROMPT)
	fmt.Println("input:", input)
}

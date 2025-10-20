package main

import (
	"encoding/json"
	"fmt"
	"time"
	"errors"
)

const (
	PROMPT = "tasks-$ "
	VERSION = "0.0.1"
)


func main() {
	
}

type Task struct {
	id int
	desc string
	status string
	createdAt time.Time
	updatedAt time.Time
}



func greeting() {
	fmt.Printf("Task Tracker v%v\n",VERSION)
}

// function fot safety input and with prompt for tasks
func smartInput(prompt ...string) (string,string,string, error) {
	if len(prompt) == 0 {
		fmt.Print("input-$ ")
	} else {
	fmt.Print(prompt[0])
	}
	var p1,p2,p3 string
	_, e := fmt.Scanln(&p1,&p2,&p3)
	if e != nil {
		e = errors.New("input error")
		return "","","",e
	}
	return p1,p2,p3,nil
}


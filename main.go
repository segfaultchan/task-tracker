package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
	"os/exec"
	"strconv"
)

const (
	PROMPT = "tasks-$ "
	VERSION = "0.1.0"
)

func main() {
	greeting()
	filename,isEmpty := getArg()
	var tasks []Task
	if !isEmpty {
		tasks,_ = importJson(filename)
	} else {
		tasks = make([]Task, 0)
	}
	run := true
	for run{
		input,_ := smartInput(PROMPT)

		switch input[0] {
		// it may not work on windows (i dont care, i write it for linux)
			case "save","s":
				if !isEmpty {
					exportJson(tasks,filename)
				} else {
					if len(input) < 2 {
						fmt.Println("needs a filename for save tasks")
						continue
					}
					exportJson(tasks,input[1])
				}
			case "clear":
				clr := exec.Command("clear")
				clr.Stdout = os.Stdout
				clr.Run()
			case "exit":
				run = false
			case "help","h":
				helpT()
			case "list","ls":
				listT(tasks)
			case "create","c","mk":
				if len(input) < 3 {
					fmt.Println("no/not enough arguments")
					continue
				}
				createT(&tasks, input[1],input[2])
			case "delete","del","d":
				if len(input) < 2 {
					fmt.Println("no/not enough arguments")
					continue
				}
				arg_int,err := strconv.Atoi(input[1])
				if (err != nil) {
					fmt.Println("error, not int argument")
					continue
				}
				deleteT(&tasks, arg_int)
			case "update","u","upd":
				if len(input) < 3 {
					fmt.Println("no/not enough arguments")
					continue
				}
				arg_int,err := strconv.Atoi(input[1])
				if (err != nil) {
					fmt.Println("error, not int argument")
					continue
				}
				updateT(&tasks,arg_int,input[2])
				fmt.Println("status updated")
			case "print","p":
				if len(input) < 2 {
					fmt.Println("no/not enough arguments")
					continue
				}
				arg_int,err := strconv.Atoi(input[1])
				if (err != nil) {
					fmt.Println("error, not int argument")
					continue
				}
				printT(tasks,arg_int)
			case "":
		case "shef","shefos","linux320":
			fmt.Println("linux 320kg shefos")
				continue
			default:
				fmt.Println("unknown command")
		}
	}
}

type Task struct {
	Desc string
	Status string
	CreatedAt string
	UpdatedAt string
}

// export will delete existing file
func exportJson(tasks []Task, fName string) error {
	os.Create(fName)
// encoding to bytes for json
	tasks_enc,err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(fName, tasks_enc, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func importJson(fName string) (tasks []Task, err error) {
	if _,err = os.Stat(fName); err != nil {
		fmt.Println(err)
		return
	}
	tasks_enc,err := os.ReadFile(fName)
	err = json.Unmarshal(tasks_enc,&tasks)
	return
}
func getArg() (arg string, isEmpty bool) {
	if len(os.Args) < 2 {
		arg = ""
		isEmpty = true
		return
	}
	arg = os.Args[1]
	isEmpty = false
	return
}

func timeToStr(t time.Time) (ts string) {
	hI,mI,sI := t.Clock()
	h,m,s := strconv.Itoa(hI),strconv.Itoa(mI),strconv.Itoa(sI)
	if hI < 10 {
		ts+="0"
		ts+=h
	} else {
		ts+=h
	}
	ts+=":"
	if mI < 10 {
		ts+="0"
		ts+=m
	} else {
		ts+=m
	}
	ts+=":"
	if sI < 10 {
		ts+="0"
		ts+=s
	} else {
		ts+=s
	}
	return ts
}

func createT (tasks *[]Task, desc string, status string) (error) {
	var task Task
	task.Desc = desc
	task.Status = status
	task.CreatedAt = timeToStr(time.Now())
	task.UpdatedAt = task.CreatedAt
	*tasks = append(*tasks, task)
	fmt.Println("task has created")
	return nil
}

func listT (tasks []Task) (error) {
	if len(tasks) == 0 {
		fmt.Println("no tasks there")
		return nil
	}
	for i := range len(tasks) {
		printT(tasks,i)
	}
	return nil
}

func deleteT (tasks *[]Task, index int) error {
	if len(*tasks) < index+1 || index<0 {
		fmt.Println("not existing index of slice")
		return nil
	}
	same := *tasks
	*tasks = append(same[:index], same[index+1:]...)
	fmt.Printf("%v element has deleted\n", index)
	return nil
}

func updateT (tasks *[]Task, index int, status string) error {
	if len(*tasks) < index+1 || index<0 {
		fmt.Println("not existing index of slice")
		return nil
	}
	(*tasks)[index].Status = status
	(*tasks)[index].UpdatedAt = timeToStr(time.Now())
	return nil
}

func printT (tasks []Task, index int) error {
	if len(tasks) < index+1 || index<0 {
		fmt.Println("not existing index of slice")
		return nil
	}
	t := tasks[index]
	fmt.Println("task id:", index)
	fmt.Println("{")
	fmt.Printf("\tdescription:\t%v\n", t.Desc )
	fmt.Printf("\tstatus:\t\t%v\n", t.Status )
	fmt.Printf("\tcreated at:\t%v\n", t.CreatedAt )
	fmt.Printf("\tupdated at:\t%v\n", t.UpdatedAt )
	fmt.Println("}")
	return nil
}

func greeting() {
	fmt.Printf("Task Tracker v%v\n",VERSION)
}

func helpT() {
	fmt.Println("help:")
	fmt.Println("task contain: id, description and status")
	fmt.Println("\t'list' - list of tasks")
	fmt.Println("\t'create !description !status' - create task")
	fmt.Println("\t'update !index !status' - update task")
	fmt.Println("\t'delete !index' - delete task")
	fmt.Println("\t'print !index' - print one task by index")
	fmt.Println("\t'save' - save tasks to file")
	fmt.Println("\t'exit' - exit program")
	fmt.Println("DONT FORGET '!'")
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
	output = strings.Split(input, " !")

	return output,nil
}

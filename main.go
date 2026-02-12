package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		printHelp()
		os.Exit(1)
	}

	var tasks []Task

	savePath, err := getSavePath()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	LoadTasks(savePath, &tasks)

	handleOperation(args, &tasks)
}

func getSavePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, err
	}

	return homeDir + "/.task.json", nil
}

func checkArgCount(args []string) error {
	if len(args) < 3 {
		return errors.New("too few arguments")
	}

	return nil
}

func printErr(err error) {
	fmt.Printf("Error: %s\n", err)
}

func printHelp() {
	fmt.Printf(`help			show this message
add "<task-body>"	add a task
del <task-id>		delete a task
done <task-id>		toggle task completion
list 			list all tasks
`)
}

func handleOperation(args []string, tasks *[]Task) {
	operation := args[1]
	switch operation {
	case "add":
		handleAdd(args, tasks)
	case "del":
		handleDelete(args, tasks)
	case "done":
		handleToggle(args, tasks)
	case "clear":
		handleClear(tasks)
	case "list":
		err := ListTasks(*tasks)
		if err != nil {
			printErr(err)
		}
	default:
		printHelp()
	}
}

func handleAdd(args []string, tasks *[]Task) {
	err := checkArgCount(args)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	input := args[2]
	*tasks = append(*tasks, Task{input, false})
	saveToFile(tasks)
}

func handleDelete(args []string, tasks *[]Task) {
	err := checkArgCount(args)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	index, err := strconv.Atoi(args[2])
	if err != nil {
		index = -1
	}
	newList, err := DeleteTask(*tasks, index-1)
	if err != nil {
		printErr(err)
	} else {
		*tasks = newList
		saveToFile(tasks)
	}
}

func handleToggle(args []string, tasks *[]Task) {
	err := checkArgCount(args)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	index, err := strconv.Atoi(args[2])
	if err != nil {
		index = -1
	}
	newTasks, err := ToggleTask(*tasks, index-1)
	if err != nil {
		printErr(err)
	} else {
		*tasks = newTasks
		saveToFile(tasks)
	}
}

func handleClear(tasks *[]Task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Clear task file? (y/N): ")
	scanner.Scan()

	switch scanner.Text() {
	case "Y", "y":
		fmt.Println("Task file cleared\n")
		*tasks = []Task{}
		saveToFile(tasks)
	default:
		os.Exit(1)
	}
}

func saveToFile(tasks *[]Task) {
	savePath, err := getSavePath()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	SaveTasks(savePath, *tasks)
}

package main

// 🌟✨ Welcome to the emoji-fied `main.go` ✨🌟
// This file orchestrates the tiny to-do CLI app — enjoy the emojis!
// 🚀🐣⚡️🎉🔥💡🛠️📦

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Operation int

const (
	ADD = iota
	EDIT
	DEL
	TOGGLE
	LIST
	CLEAR
)

// main is the entrypoint for the CLI 🏁
// Usage: `task <command> [arguments]` 🧭
func main() {
	args := os.Args

	if len(args) <= 1 {
		// No args -> show help 🤔📘
		printHelp() // 📚
		os.Exit(1) // ❌
	}

	// load tasks into memory 🧠📥
	var tasks []Task // 🗂️ list of tasks

	savePath, err := getSavePath() // 🔍 find save location
	if err != nil {
		printErr(err) // 🆘 show error
		os.Exit(1)
	}

	// Load previously saved tasks from disk 📂➡️🧠
	LoadTasks(savePath, &tasks)

	// Dispatch the requested operation (add/edit/list/etc.) 🏷️
	handleOperation(args, &tasks)
}

func getSavePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, err
	}

	// default save file is in user's home directory 🏠
	return homeDir + "/.task.json", nil // 💾
}

func checkArgCount(args []string, op Operation) error {
	switch op {
	case ADD, DEL, TOGGLE:
		if len(args) < 3 {
			// not enough args for add/del/toggle 😅
			return errors.New("too few arguments")
		}
	case EDIT:
		if len(args) < 4 {
			// edit needs id + new text ✍️
			return errors.New("too few arguments")
		}
	default:
		// unknown operation 🧨
		return errors.New("invalid operation")
	}

	return nil
}

func printErr(err error) {
	// print errors in a friendly format ⚠️
	fmt.Printf("Error: %s\n", err)
}

func printHelp() {
	// help text with emoji hints 🆘
	fmt.Printf(`
task <command> [arguments]

help			show this message 📖❓
add [text]		add a task ➕📝
edit <id> <text>	edit a task ✏️
del [id]		delete a task 🗑️
done [id]		toggle task completion ✅/⬜️
list 			list all tasks 📜

`) // 🎯
}

func handleOperation(args []string, tasks *[]Task) {
	operation := args[1]
	switch operation {
	case "add":
		handleAdd(args, tasks)
	case "edit":
		handleEdit(args, tasks)
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
		// unknown command -> show help 🤷‍♀️
		printHelp()
	}
}

func handleAdd(args []string, tasks *[]Task) {
	err := checkArgCount(args, ADD)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	input := args[2]
	// append new task (not done by default) ➕
	*tasks = append(*tasks, Task{input, false})
	saveToFile(tasks) // persist to disk 💾
}

func handleEdit(args []string, tasks *[]Task) {
	err := checkArgCount(args, EDIT)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	index, err := strconv.Atoi(args[2])
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	text := args[3]
	*tasks, err = EditTask(*tasks, index-1, text)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	// save after editing ✨
	saveToFile(tasks)
}

func handleDelete(args []string, tasks *[]Task) {
	err := checkArgCount(args, DEL)
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
		// persist deletion 🧹
		saveToFile(tasks)
	}
}

func handleToggle(args []string, tasks *[]Task) {
	err := checkArgCount(args, TOGGLE)
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
		// toggle complete/incomplete 🔁
		saveToFile(tasks)
	}
}

func handleClear(tasks *[]Task) {
	scanner := bufio.NewScanner(os.Stdin)
	// confirm destructive action 🛑
	fmt.Print("Clear task file? (y/N): ")
	scanner.Scan()

	switch scanner.Text() {
	case "Y", "y":
		// user confirmed -> clear list 🧽
		fmt.Println("Task file cleared")
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
	// write tasks to JSON file on disk 📝➡️💾
	SaveTasks(savePath, *tasks)
}

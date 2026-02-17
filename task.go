package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// 🗂️ Task represents a single to-do item with text and completion state
type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// ToggleTask flips the Done state for a task at index ⚖️
// returns an error if the index is invalid 🚫
func ToggleTask(taskList []Task, index int) ([]Task, error) {
	if index >= len(taskList) || index < 0 {
		return taskList, errors.New("invalid index") // ❌
	}

	// flip the boolean ✅ -> ⬜️ or ⬜️ -> ✅
	taskList[index].Done = !taskList[index].Done
	return taskList, nil

}

// EditTask updates the Text field for the task at index ✍️
func EditTask(taskList []Task, index int, text string) ([]Task, error) {
	if index >= len(taskList) || index < 0 {
		return taskList, errors.New("invalid index")
	}

	taskList[index].Text = text // 📝
	return taskList, nil
}

// DeleteTask removes the task at index and returns a new slice 🧹
func DeleteTask(taskList []Task, index int) ([]Task, error) {
	if index >= len(taskList) || index < 0 {
		return taskList, errors.New("invalid index")
	}

	newList := []Task{}

	// copy all tasks except the one being deleted 🔁
	for i, task := range taskList {
		if i != index {
			newList = append(newList, task)
		}
	}

	return newList, nil
}

// LoadTasks reads JSON from fileName and unmarshals into taskList 📂➡️🔧
// If the file doesn't exist it simply returns silently (first-run friendly) 🌱
func LoadTasks(fileName string, taskList *[]Task) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return // no file yet, that's fine 👍
		}
	} else {
		json.Unmarshal(data, taskList) // best-effort unmarshal 🧪
	}
}

// SaveTasks marshals the tasks and writes them to fileName 💾
func SaveTasks(fileName string, taskList []Task) {
	data, err := json.Marshal(taskList)
	if err != nil {
		printErr(err) // inform user of marshal failure 😬
	} else {
		err = os.WriteFile(fileName, data, 0644) // write with safe perms 🔐
		if err != nil {
			printErr(err) // inform user of write failure 🆘
		}
	}
}

// ListTasks prints tasks to stdout, returning an error if empty 📜
func ListTasks(taskList []Task) error {
	if len(taskList) == 0 {
		return errors.New("task list empty") // nothing to show 😴
	}
	for i, task := range taskList {
		if task.Done {
			fmt.Print("[X] ") // done ✅
		} else {
			fmt.Print("[ ] ") // not done ⬜️
		}

		fmt.Printf("%d. %s\n", i+1, task.Text) // nice numbered list 🔢
	}

	return nil
}

package main

import (
	"fmt"
	"learning_go/internal/taskmanager"
)

func main() {
	var taskList []string = make([]string, 0)
	var taskStatus = make(map[string]bool)
	var enterTask string
	var command string
	var status string

	fmt.Println("🗂️ WELCOME TO Ei4 TASK MANAGER")
	fmt.Println("- This program will help you keep tab of your incomplete and completed task.")
TaskLoop:
	for {
		// ENTER ANOTHER TASK (LOOP)
		taskmanager.DisplayTask(taskList, status, taskStatus)
		command = taskmanager.CommandSection(command)

		switch command {
		case "0":
			taskList, status = taskmanager.InsertTask(taskList, taskStatus, enterTask, status)
			continue
		case "1":
			taskmanager.MarkComplete(taskList, taskStatus)
			continue
		case "2":
			taskmanager.MarkIncomplete(taskList, taskStatus)
			continue
		case "3":
			taskmanager.EditTask(taskList)
			continue
		case "4":
			taskList, taskStatus = taskmanager.RemoveTask(taskList, taskStatus)
			continue
		case "X":
			fmt.Println("🗂️ Exiting Program.")
			break TaskLoop
		default:
			fmt.Println("❗ Error has occured.")
			break TaskLoop
		}
	}
}

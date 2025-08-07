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

	fmt.Println("📋 WELCOME TO Ei4 TASK MANAGER")
	fmt.Println("- This program will help you keep tab of your incomplete and completed task.")
TaskLoop:
	for {
		// ENTER ANOTHER TASK (LOOP)
		taskmanager.DisplayTask(taskList, status, taskStatus)
		command = taskmanager.CommandSection(command)

		if command == "0" {
			taskList, status = taskmanager.InsertTask(taskList, taskStatus, enterTask, status)
			continue
		} else if command == "1" {
			taskmanager.MarkComplete(taskList, taskStatus)
			continue
		} else if command == "2" {
			taskmanager.MarkIncomplete(taskList, taskStatus)
		} else if command == "3" {

		} else if command == "X" {
			fmt.Println("📋 Exiting Program.")
			break TaskLoop
		} else {
			fmt.Println("❗ Error has occured.")
			break TaskLoop
		}
	}
}

package taskmanager

import (
	"fmt"
	"strconv"
)

func RemoveTask(taskList []string, taskStatus map[string]bool) ([]string, map[string]bool) {
	var num string
	var taskToRemove string
	var command string

	for {
		fmt.Print("✏️ Enter the task number to remove: ")
		fmt.Scanln(&num)

		var index, err = strconv.Atoi(num)
		if err != nil || index < 1 || index > len(taskList) {
			fmt.Println("❗ Invalid task number, please enter the correct task number.")
			continue
		}
		taskToRemove = taskList[index-1]

	CommandLoop:
		for {
			fmt.Printf("📝 You're removing \"%d. %s\" task.\n", index, taskList[index-1])
			fmt.Print("❗ Are you sure? (Y/N): ")
			fmt.Scanln(&command)

			switch command {
			case "Y":
				// REMOVING TASK FROM SLICE (MAKING A NEW SLICE)
				taskList = append(taskList[:index-1], taskList[index:]...)
				// REMOVING TASK FROM MAP
				delete(taskStatus, taskToRemove)
				break CommandLoop
			case "N":
				fmt.Println("📝 Going back to Command Line Interface.")
				break CommandLoop
			default:
				fmt.Println("❗ Invalid command, please enter a correct command.")
				continue
			}
		}
		return taskList, taskStatus
	}
}

package taskmanager

import (
	"fmt"
	"learning_go/internal/inputvalidator"
)

func RemoveTask(taskList []string, taskStatus map[string]bool) ([]string, map[string]bool) {
	var num string
	var taskToRemove string
	var command string

CommandLoop:
	for {
		updatedList, ok := inputvalidator.TaskAvailability(taskList)
		if !ok {
			break
		}
		taskList = updatedList

		fmt.Print("✏️ Enter the task number to remove: ")
		fmt.Scanln(&num)

		var index, err = inputvalidator.IsValidNumberInput(num, 1, len(taskList), "task number")
		if !err {
			break
		}

		taskToRemove = taskList[index-1]

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
				fmt.Println("📝 Returning to Command Line Interface.")
				break CommandLoop
			default:
				fmt.Println("❗ Invalid command, please enter a correct command.")
				continue
			}
		}
	}
	return taskList, taskStatus
}

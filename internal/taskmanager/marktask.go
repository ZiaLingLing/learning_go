package taskmanager

import (
	"fmt"
	"learning_go/internal/inputvalidator"
)

func MarkComplete(taskList []string, taskStatus map[string]bool) {
	var num string
	var task string
	for {
		var taskList, ok = inputvalidator.TaskAvailability(taskList)
		if !ok {
			break
		}

		fmt.Print("✏️ Enter the task number to mark as completed: ")
		fmt.Scanln(&num)

		var index, err = inputvalidator.IsValidNumberInput(num, 0, len(taskList), "task number")
		if !err {
			continue
		}

		task = taskList[index-1]

		if taskStatus[task] {
			fmt.Println("❗ Task already completed.")
			break
		}

		taskStatus[task] = true
		fmt.Printf("📝 Task \"%s\" marked as completed.\n", task)
		break
	}
}

func MarkIncomplete(taskList []string, taskStatus map[string]bool) {
	var num string
	var task string

	for {
		var taskList, ok = inputvalidator.TaskAvailability(taskList)
		if !ok {
			break
		}

		fmt.Print("✏️ Enter the task number to mark as incomplete: ")
		fmt.Scanln(&num)

		var index, err = inputvalidator.IsValidNumberInput(num, 0, len(taskList), "task number")
		if !err {
			continue
		}

		task = taskList[index-1]

		if !taskStatus[task] {
			fmt.Println("❗ Task still incomplete.")
			break
		}

		taskStatus[task] = false
		fmt.Printf("📝 Task \"%s\" marked as incomplete.\n", task)
		break
	}
}

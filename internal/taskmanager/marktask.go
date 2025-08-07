package taskmanager

import (
	"fmt"
	"strconv"
)

func MarkComplete(taskList []string, taskStatus map[string]bool) {
	var mark string
	var task string

	for {
		fmt.Print("✏️ Enter the task number to mark as completed: ")
		fmt.Scanln(&mark)

		var index, err = strconv.Atoi(mark)
		if err != nil || index < 1 || index > len(taskList) {
			fmt.Println("❗ Invalid task number, please enter the correct task number.")
			continue
		}

		task = taskList[index-1]

		if taskStatus[task] {
			fmt.Println("❗ Task already completed.")
			break
		}

		taskStatus[task] = true
		fmt.Printf("✏️ Task \"%s\" marked as completed.\n", task)
		break
	}
}

func MarkIncomplete(taskList []string, taskStatus map[string]bool) {
	var mark string
	var task string

	for {
		fmt.Print("✏️ Enter the task number to mark as incomplete: ")
		fmt.Scanln(&mark)

		var index, err = strconv.Atoi(mark)
		if err != nil || index < 1 || index > len(taskList) {
			fmt.Println("❗ Invalid task number, please enter the correct task number.")
			continue
		}

		task = taskList[index-1]

		if taskStatus[task] == false {
			fmt.Println("❗ Task still incomplete.")
			break
		}

		taskStatus[task] = false
		fmt.Printf("✏️ Task \"%s\" marked as incomplete.\n", task)
		break
	}
}

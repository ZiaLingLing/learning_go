package taskmanager

import (
	"bufio"
	"fmt"
	"learning_go/internal/inputvalidator"
	"os"
)

func EditTask(taskList []string) []string {
	var num string
	var rename string

	for {
		var taskList, ok = inputvalidator.TaskAvailability(taskList)
		if !ok {
			break
		}

		fmt.Print("✏️ Enter the task number to edit: ")
		fmt.Scanln(&num)

		var index, err = inputvalidator.IsValidNumberInput(num, 0, len(taskList), "task number")
		if !err {
			continue
		}

		fmt.Printf("📝 You're editing \"%d. %s\" task.\n", index, taskList[index-1])
		fmt.Print("✏️ Rename task: ")
		var scanner = bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			rename = scanner.Text()
		}
		taskList[index-1] = rename

		fmt.Println("📝 Task renamed.")
		break
	}
	return taskList
}

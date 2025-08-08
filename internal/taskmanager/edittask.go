package taskmanager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EditTask(taskList []string) []string {
	var num string
	var rename string

	for {
		fmt.Print("✏️ Enter the task number to edit: ")
		fmt.Scanln(&num)

		var index, err = strconv.Atoi(num)
		if err != nil || index < 0 || index > len(taskList) {
			fmt.Println("❗ Invalid task number, please enter the correct task number.")
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
		return taskList
	}
}

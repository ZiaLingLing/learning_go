package taskmanager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EditTask(taskList []string) string {
	var index int8

	for {
		fmt.Println("✏️ Enter the task number to edit: ")
		fmt.Scanln(&index)

		var index, err = strconv.Atoi(index)
		if err != nil || index < 0 || index > len(taskList) {
			fmt.Println("❗ Invalid task number, please enter the correct task number.")
			continue
		}

		fmt.Printf("✏️ You're editing \"%d. %s\" task.")
		fmt.Print("✏️ Rename task: ")
		var scanner = bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			rename = scanner.Text()
		}

		taskList[index] = ""
	}
}

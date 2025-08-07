package taskmanager

import (
	"bufio"
	"fmt"
	"os"
)

func InsertTask(taskList []string, taskStatus map[string]bool, enterTask string, status string) ([]string, string) {
	// INSERTING TASK TO MAP
	fmt.Print("✏️ Enter a new task: ")
	var scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		enterTask = scanner.Text()
	}
	taskList = append(taskList, enterTask)
	fmt.Printf("✏️ New task \"%s\" has been added.", enterTask)

	// SETTING TASK STATUS
	taskStatus[enterTask] = false
	switch taskStatus[enterTask] {
	case false:
		status = "❌"
	case true:
		status = "✅"
	}

	return taskList, status
}

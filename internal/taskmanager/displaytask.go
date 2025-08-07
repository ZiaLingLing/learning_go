package taskmanager

import "fmt"

func DisplayTask(taskList []string, status string, taskStatus map[string]bool) {

	if len(taskList) < 1 {
		fmt.Println("\n📝 PENDING TASKS: ")
		fmt.Println("🎉 You have no pending task at the moment.")
		return
	}
	fmt.Println("\n📝 PENDING TASKS: ")

	for index, task := range taskList {
		switch taskStatus[task] {
		case false:
			status = "❌"
		case true:
			status = "✅"
		}
		fmt.Printf("%d. %v: %s\n", index+1, task, status)
	}
}

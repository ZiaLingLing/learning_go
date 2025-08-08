package taskmanager

import (
	"fmt"
	"learning_go/internal/inputvalidator"
)

func CommandSection(command string) string {
	fmt.Print("\n", `🖥️ COMMAND LINE INTERFACE:
Enter [0] to insert a new task.
Enter [1] to mark task completed.
Enter [2] to mark task incomplete.
Enter [3] to edit task.
Enter [4] to delete task.
Enter [X] to exit program.`, "\n")

	for {
		fmt.Print("🖥️ ENTER COMMAND: ")
		fmt.Scanln(&command)

		if command == "X" {
			return command
		}
		var _, err = inputvalidator.IsValidNumberInput(command, 0, 4, "command")
		if !err {
			continue
		}
		return command
	}
}

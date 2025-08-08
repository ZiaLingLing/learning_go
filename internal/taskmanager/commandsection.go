package taskmanager

import (
	"fmt"
	"strconv"
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
		var num, err = strconv.Atoi(command)
		if err != nil || 0 > num || num > 4 {
			fmt.Println("❗ Invalid command, please enter a correct command.")
			continue
		}
		return command
	}
}

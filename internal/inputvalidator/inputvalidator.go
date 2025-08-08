package inputvalidator

import (
	"fmt"
	"strconv"
)

func TaskAvailability(taskList []string) ([]string, bool) {
	if len(taskList) == 0 {
		fmt.Println("❗ No tasks to remove.")
		fmt.Println("🖥️ Returning to Command Line Interface.")
		return taskList, false
	}
	return taskList, true
}

func IsValidNumberInput(input string, min int, max int, object string) (int, bool) {
	var index, err = strconv.Atoi(input)
	if err != nil || index < min || index > max {
		fmt.Printf("❗ Invalid %s, please enter the correct %s.\n", object, object)
		return 0, false
	}
	return index, true
}

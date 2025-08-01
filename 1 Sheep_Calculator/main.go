package main

import (
	"errors"
	"fmt"
)

// SCENARIO: ANTHONY IS YOUR FRIEND, HE WANTS TO BUY YOUR SHEEP BUT HE ONLY WANTS 10 MULTIPLES OF YOUR SHEEP
// THIS IS A SILLY PROGRAM THAT HELPS YOU COUNT YOUR SHEEP AND HOW MANY YOU CAN SELL TO ANTHONY

func main() {
	var sheepOwned int32 = 90
	var sheepNew int32 = 189
	var err error

	fmt.Printf("You have %v sheep.\n", sheepOwned)
	fmt.Printf("You are adding %v sheep to your sheep pen.\n", sheepNew)
	sheepOwned, err = addSheep(sheepOwned, sheepNew)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("You now have %v sheep.\n", sheepOwned)
	}

	// CHECKING AMOUNT OF SHEEP AND ANTHONY'S RESPONSE
	var multiple, remainder = sellingSheep(sheepOwned)

	if multiple < 1 {
		fmt.Println("Total sheep aren't enough to sell.")
	} else if multiple == 1 {
		fmt.Printf("You can sell %v group of 10 sheep to Anthony with %v remaining.", multiple, remainder)
	} else if multiple > 1 {
		fmt.Printf("You can sell %v groups of 10 sheep to Anthony with %v sheep remaining.", multiple, remainder)
	}
}

func addSheep(sheepOwned int32, sheepNew int32) (int32, error) {
	// ERROR CHECK
	var err error
	if sheepOwned < 0 {
		err = errors.New("[ERROR] Invalid number of sheep owned")
		return 0, err
	} else if sheepNew <= 0 {
		err = errors.New("[ERROR] Invalid number of added sheep")
		return 0, err
	}

	// ADDITION
	sheepOwned = sheepOwned + sheepNew
	return sheepOwned, err
}

func sellingSheep(sheepOwned int32) (int32, int32) {
	var multiple int32 = sheepOwned / 10
	var remainder int32 = sheepOwned % 10

	return multiple, remainder
}

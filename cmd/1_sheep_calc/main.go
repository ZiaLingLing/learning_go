package main

import (
	"fmt"

	"learning_go/internal/sheepcalc"
)

// SCENARIO: ANTHONY IS YOUR FRIEND, HE WANTS TO BUY YOUR SHEEP BUT HE ONLY WANTS 10 MULTIPLES OF YOUR SHEEP
// THIS IS A SILLY PROGRAM THAT HELPS YOU COUNT YOUR SHEEP AND HOW MANY YOU CAN SELL TO ANTHONY (No user input)

func main() {
	var sheepOwned int32 = 90
	var sheepNew int32 = 189
	var err error

	fmt.Println("Scenario: Anthony is your friend, he wants to buy your sheep under the requirement that they must be in multiples of 10.")

	fmt.Printf("You have %v sheep.\n", sheepOwned)
	fmt.Printf("You are adding %v sheep to your sheep pen.\n", sheepNew)
	sheepOwned, err = sheepcalc.AddSheep(sheepOwned, sheepNew)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("You now have %v sheep.\n", sheepOwned)
	}

	// CHECKING AMOUNT OF SHEEP AND ANTHONY'S RESPONSE
	var multiple, remainder = sheepcalc.SellingSheep(sheepOwned)

	if multiple < 1 {
		fmt.Println("Total sheep aren't enough to sell.")
	} else if multiple == 1 {
		fmt.Printf("You can sell %v group of 10 sheep to Anthony with %v remaining.", multiple, remainder)
	} else if multiple > 1 {
		fmt.Printf("You can sell %v groups of 10 sheep to Anthony with %v sheep remaining.", multiple, remainder)
	}
}

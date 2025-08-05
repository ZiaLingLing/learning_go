package sheepcalc

import (
	"errors"
)

func AddSheep(sheepOwned int32, sheepNew int32) (int32, error) {
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

func SellingSheep(sheepOwned int32) (int32, int32) {
	var multiple int32 = sheepOwned / 10
	var remainder int32 = sheepOwned % 10

	return multiple, remainder
}

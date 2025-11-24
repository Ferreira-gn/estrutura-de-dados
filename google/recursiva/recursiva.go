package recursiva

import (
	"fmt"
	"time"
)

func SolvedInRecursiveMethod(initPoint int) {
	var biggestSequency int

	initTime := time.Now()

	for initPoint > 1 {
		currentSequency := calculateLongestSequency(initPoint)

		if currentSequency > biggestSequency {
			biggestSequency = currentSequency
			continue
		}

		initPoint--
	}

	timeDuration := time.Since(initTime)

	fmt.Printf("\n\nMetódo incrmental")
	fmt.Printf("\nMaior sequencia %-34s tempo \n", "")
	fmt.Printf("%-50d %fs", biggestSequency, timeDuration.Seconds())
}

func calculateLongestSequency(initPoint int) int {

	if (initPoint == 1){
		return 0
	}

	if initPoint%2 != 0 {
		count := calculateLongestSequency(3*initPoint + 1)
		return count + 1
	}

	count := calculateLongestSequency(initPoint / 2)
	return  count+1
}

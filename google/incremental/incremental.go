package incremental

import (
	"fmt"
	"time"
)

func SolvedInIncrementalMethod(initPoint int) {
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

	fmt.Printf("\n\nMetódo recursivo")
	fmt.Printf("\nMaior sequencia %-34s tempo \n", "")
	fmt.Printf("%-50d %fs", biggestSequency, timeDuration.Seconds())
}

func calculateLongestSequency(initPoint int) int {
	var count int = 0

	for initPoint != 1 {
		if initPoint%2 != 0 {
			initPoint = 3*initPoint + 1
			count++
			continue
		}

		initPoint = initPoint / 2
		count++
	}

	return count
}

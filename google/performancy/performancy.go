package performancy

import (
	"fmt"
	"time"
)

func SolvedInPreformatMethod(initPoint int) {
	var biggestSequency int

	hashMap := make([]int, initPoint+1)

	initTime := time.Now()

	for i := 1; i < initPoint; i++ {
		currentSequency := calculateLongestSequency(i, hashMap)

		if currentSequency > biggestSequency {
			biggestSequency = currentSequency
			continue
		}
	}

	timeDuration := time.Since(initTime)

	fmt.Printf("\n\nMetódo performatico")
	fmt.Printf("\nMaior sequencia %-34s tempo \n", "")
	fmt.Printf("%-50d %fs", biggestSequency, timeDuration.Seconds())
}

func calculateLongestSequency(initPoint int, hashMap []int) int {

	if hashMap[initPoint] != 0 {
		sequencyAlreadyCoverage := hashMap[initPoint]

		return sequencyAlreadyCoverage
	}

	count := 0
	currentPoint := initPoint

	for currentPoint > 1 {
		if currentPoint < len(hashMap) && hashMap[currentPoint] != 0 {
			alreadyCoverage := hashMap[currentPoint]
			count = count + alreadyCoverage
			break
		}

		if currentPoint%2 != 0 {
			currentPoint = 3*currentPoint + 1
			count++
			continue
		}

		currentPoint = currentPoint / 2
		count++
	}

	if initPoint < len(hashMap) {
		hashMap[initPoint] = count
	}

	return count
}

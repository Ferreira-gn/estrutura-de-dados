package hashmap

import (
	"fmt"
	"time"
)

func SolvedInPreformatMethod(initPoint int) {
	var biggestSequency int 
	hashMap := make(map[int]int)
	initTime := time.Now()

	for i := 1; i < initPoint; i++ {
		currentSequency := calculateLongestSequency(i, hashMap)

		if currentSequency > biggestSequency {
			biggestSequency = currentSequency
			continue
		}
	}

	timeDuration := time.Since(initTime)

	fmt.Printf("\n\nMetódo com a estrutura hashmap do go")
	fmt.Printf("\nMaior sequencia %-34s tempo \n", "")
	fmt.Printf("%-50d %fs", biggestSequency, timeDuration.Seconds())
}

func calculateLongestSequency(initPoint int, hashMap map[int]int) int {
	sequencyAlreadyCoverage, includeInHashMap := hashMap[initPoint]
	
	if includeInHashMap {
		return  sequencyAlreadyCoverage
	}

	count := 0
	currentPoint := initPoint

	for currentPoint > 1 {
		alreadyCoverage , includeInHashMap := hashMap[currentPoint]

		if includeInHashMap {
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

	hashMap[initPoint] = count
	return count
}
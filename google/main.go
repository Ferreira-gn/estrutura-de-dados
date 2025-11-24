package main

import (
	hashmap "github.com/Dom-Garotom/estruturaDeDados/google-challenger/hash"
	"github.com/Dom-Garotom/estruturaDeDados/google-challenger/incremental"
	"github.com/Dom-Garotom/estruturaDeDados/google-challenger/performancy"
	"github.com/Dom-Garotom/estruturaDeDados/google-challenger/recursiva"
)

func main(){
	initPoint := 10000000
	
	
	incremental.SolvedInIncrementalMethod(initPoint)
	recursiva.SolvedInRecursiveMethod(initPoint)
	hashmap.SolvedInPreformatMethod(initPoint)
	performancy.SolvedInPreformatMethod(initPoint)
}

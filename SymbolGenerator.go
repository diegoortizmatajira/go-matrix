package main

import "math/rand"

type SymbolRange struct {
	name  string
	start int
	end   int
}

func generateSymbol() rune {
	symbolRanges := []SymbolRange{
		{
			name:  "Normal",
			start: 33,
			end:   126,
		},
		{
			name:  "Japanese",
			start: 12448,
			end:   12538,
		},
		{
			name:  "Cyrilic",
			start: 1040,
			end:   1103,
		},
	}
	symbolCount := 0
	for _, v := range symbolRanges {
		symbolCount += v.end - v.start
	}
	randomNumber := rand.Intn(symbolCount)
	symbolCount = 0
	for _, v := range symbolRanges {
		rangeCount := v.end - v.start
		localIndex := randomNumber - symbolCount
		if localIndex < rangeCount {
			return rune(v.start + localIndex)
		}
		symbolCount += rangeCount
	}
	return '#'
}

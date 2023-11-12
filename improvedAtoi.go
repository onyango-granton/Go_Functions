package main

import "fmt"

func isValidInt(s string) bool {
	for _, ch := range s {
		if !(ch >= '0' && ch <= '9' || ch == '+' || ch == '-') {
			return false
		}
	}
	return true
}

func stringToInt(s string) int {
	finalString := ""
	result := 0
	for _, ch := range s {
		if ch == '+' || ch == '-' {
			continue
		}
		finalString = finalString + string(ch)
	}
	for _, ch := range finalString {
		result = result*10 + int(ch-'0')
	}
	return result
}

func finalSign(s string) int {
	result := 1
	positiveResult := 1
	negativeResult := 1
	negativeCount := 0
	positiveCount := 0
	for _, ch := range s {
		if ch == '-' {
			negativeCount++
		}
		if ch == '+' {
			positiveCount++
		}
	}
	if positiveCount >= 1 && negativeCount >= 1 {
		for i := 0; i < positiveCount; i++ {
			positiveResult = positiveResult * 1
		}
		for i := 0; i < negativeCount; i++ {
			negativeResult = negativeResult * -1
		}
		result = positiveResult * negativeResult
	} else if positiveCount >= 1 {
		for i := 0; i < positiveCount; i++ {
			result = result * 1
		}
	} else if negativeCount >= 1 {
		for i := 0; i < negativeCount; i++ {
			result = result * -1
		}
	}
	return result
}

func intToString(num int) string {
	result := ""
	i := '0'
	if num == 0 {
		return "0"
	}
	for j := 0; j < num%10; j++ {
		i++
	}
	if num/10 != 0 {
		// intToString(num / 10)
		result = intToString(num/10) + string(i)
	}
	if num/10 == 0 {
		result = result + string(i)
	}

	return result
}

func Atoi(s string) int {
	result := 0
	if isValidInt(s) {
		result = stringToInt(s) * finalSign(s)
	}

	return result
}

func main() {
	s1 := "Hello world"
	s2 := "-5644344"
	s3 := "++452336"
	s4 := "-3553234"
	s5 := "-56344-323"
	s6 := "4534t345"

	// i1 := 3552353532

	var testSlice []string
	testSlice = []string{s1, s2, s3, s4, s5, s6}

	for _, s := range testSlice {
		fmt.Println(Atoi(s))
	}
	// fmt.Println(intToString(i1))
}

package day_one

import (
	"unicode"
)

func ExtractCalibrationValue(line string) int {
	var firstValue, lastValue rune

	for _, char := range line {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			firstValue = char
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) || unicode.IsLetter(rune(line[i])) {
			lastValue = rune(line[i])
			break
		}
	}

	if firstValue == 0 || lastValue == 0 {
		return 0
	}

	res := int(firstValue-'0')*10 + int(lastValue-'0')
	println(res)
	return res
}

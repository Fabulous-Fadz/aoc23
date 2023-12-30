package trebuchet

import (
	"github.com/Fabulous-Fadz/aoc23/day1/trebuchet/word"
	"github.com/Fabulous-Fadz/aoc23/internal"
)

var (
	numbers = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	words = word.New()
)

const baseNumber = -1

func init() {
	for key := range numbers {
		words.Insert(key)
		words.InsertReverse(key)
	}
}

func DecodeInput(code []byte) int {
	left, right := 0, len(code)-1

	for {
		leftOk, rightOk := internal.IsDigit(code[left]), internal.IsDigit(code[right])

		if leftOk && rightOk {
			break
		}

		if !leftOk {
			left++
		}
		if !rightOk {
			right--
		}
	}

	return 10*int(code[left]-'0') + int(code[right]-'0')
}

func DecodeWithWords(code []byte) int {
	left, right := 0, len(code)-1
	leftNum, rightNum := baseNumber, baseNumber

	for left < len(code) || right >= 0 {
		leftOk, rightOk := leftNum != baseNumber, rightNum != baseNumber

		if leftOk && rightOk {
			break
		}

		// Now check left for right numbers...
		var i int
		for i = left; !leftOk && i < len(code) && !words.ContainsPrefix(string(code[i])); i++ {
			if internal.IsDigit(code[i]) {
				leftOk = true
				leftNum = int(code[i] - '0')
			}
		}

		if !leftOk {
			left = i
		}

		// Here either leftOk==true or we have found a prefix
		for i = left + 2; !leftOk && i < len(code) && words.ContainsPrefix(string(code[left:i+1])); i++ {
			if words.Contains(string(code[left : i+1])) {
				leftOk = true
				leftNum = numbers[string(code[left:i+1])]
			}
		}

		for ; !rightOk && right >= 0 && !words.ContainsSuffix(string(code[right])); right-- {
			if internal.IsDigit(code[right]) {
				rightOk = true
				rightNum = int(code[right] - '0')
			}
		}

		// Here either rightOk == true or we have found a suffix...
		for i = right - 2; !rightOk && i >= 0 && words.ContainsSuffix(string(code[i:right+1])); i-- {
			if words.Contains(string(code[i : right+1])) {
				rightOk = true
				rightNum = numbers[string(code[i:right+1])]
			}
		}

		left++
		right--
	}

	return 10*leftNum + rightNum
}

func DecodeAll(codes [][]byte) (sum int) {
	for _, code := range codes {
		sum += DecodeInput(code)
	}

	return
}

func DecodeAllWords(codes [][]byte) (sum int) {
	for _, code := range codes {
		sum += DecodeWithWords(code)
	}

	return
}

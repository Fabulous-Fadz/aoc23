package trebuchet

import "github.com/Fabulous-Fadz/aoc23/day1/trebuchet/word"

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
		leftOk, rightOk := isDigit(code[left]), isDigit(code[right])

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
		for i = left; !leftOk && i < len(code); i++ {
			if isDigit(code[left]) {
				leftNum = int(code[left] - '0')
				left, leftOk = i, true
				break
			}
			// try to match once we no longer find a match...
			if words.ContainsPrefix(string(code[left : i+1])) {
				continue
			} else if words.Contains(string(code[left:i])) { // +1 removed because we're looking at previous index now.
				// If we found the word in the trie then we have it in the dictionary...
				leftOk = true
				leftNum = numbers[string(code[left:i])]
			}
		}

		// Check on the right...
		for i = right - 1; !rightOk && i >= 0; i-- {
			if isDigit(code[right]) {
				rightNum = int(code[right] - '0')
				right, rightOk = i, true
				break
			}

			if words.ContainsSuffix(string(code[i : right+1])) {
				continue
			} else if word := string(code[i+1 : right+1]); words.Contains(word) { //words.Contains(string(code[i+1:right+1])){
				rightOk = true
				rightNum = numbers[word]
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

func isDigit(char byte) bool {
	diff := char - '0'
	return diff <= 9
}

package trebuchet

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

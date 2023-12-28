package internal

func IsDigit(input byte) bool {
	res := input - '0'
	return res <= 9
}

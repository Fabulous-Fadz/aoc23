package internal

import "log"

func IsDigit(input byte) bool {
	res := input - '0'
	return res <= 9
}

func Must[T interface{}](result T, err error) T {
	if err != nil {
		log.Panicf("Could not complete operation. %v", err)
	}
	return result
}

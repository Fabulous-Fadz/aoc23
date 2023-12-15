package trebuchet_test

import (
	"testing"

	trebuchet "github.com/Fabulous-Fadz/aoc23/Day1/Trebuchet"
)

var (
	testCases = []struct {
		name  string
		input []byte
		want  int
	}{
		{"1abc2", []byte("1abc2"), 12},
		{"pqr3stu8vwx", []byte("pqr3stu8vwx"), 38},
		{"a1b2c3d4e5f", []byte("a1b2c3d4e5f"), 15},
		{"treb7uchet", []byte("treb7uchet"), 77},
	}

	allTests = []struct {
		name  string
		input [][]byte
		want  int
	}{
		{
			"TestAll",
			[][]byte{
				[]byte("1abc2"),
				[]byte("pqr3stu8vwx"),
				[]byte("a1b2c3d4e5f"),
				[]byte("treb7uchet"),
			},
			142,
		},
	}
)

func TestDecode(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			have := trebuchet.DecodeInput(tc.input)
			if tc.want != have {
				t.Errorf("Decoding failed. Want %v, have %v", tc.want, have)
			}
		})
	}
}

func TestTestAll(t *testing.T) {
	for _, tc := range allTests {
		t.Run(tc.name, func(t *testing.T) {
			have := trebuchet.DecodeAll(tc.input)
			if tc.want != have {
				t.Errorf("Decoding all failed. Want %v, have %v", tc.want, have)
			}
		})
	}
}

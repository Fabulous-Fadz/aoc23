package trebuchet_test

import (
	"testing"

	"github.com/Fabulous-Fadz/aoc23/day1/trebuchet"
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

	extraCases = []struct {
		name  string
		input []byte
		want  int
	}{
		{"two1nine", []byte("two1nine"), 29},
		{"eightwothree", []byte("eightwothree"), 83},
		{"abcone2threexyz", []byte("abcone2threexyz"), 13},
		{"xtwone3four", []byte("xtwone3four"), 24},
		{"4nineeightseven2", []byte("4nineeightseven2"), 42},
		{"zoneight234", []byte("zoneight234"), 14},
		{"7pqrstsixteen", []byte("7pqrstsixteen"), 76},
		{"88", []byte("eight7eightcx"), 88},
		{"12", []byte("1onetwo2"), 12},
		{"66", []byte("6vtg"), 66},
		{"88-dude", []byte("peightp"), 88},
		{"89-dude", []byte("peight9p"), 89},
		{"eleven", []byte("one"), 11},
		{"11-again", []byte("12jsdmvonesth"), 11},
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

	allWords = struct {
		name  string
		input [][]byte
		want  int
	}{
		"TestAllWords",
		[][]byte{
			[]byte("two1nine"),
			[]byte("eightwothree"),
			[]byte("abcone2threexyz"),
			[]byte("xtwone3four"),
			[]byte("4nineeightseven2"),
			[]byte("zoneight234"),
			[]byte("7pqrstsixteen"),
		},
		281,
	}
)

func TestDecode(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//have := trebuchet.DecodeInput(tc.input)
			have := trebuchet.DecodeInput(tc.input)
			if tc.want != have {
				t.Errorf("Decoding failed. Want %v, have %v", tc.want, have)
			}
		})
	}
}

func TestDecodeAll(t *testing.T) {
	for _, tc := range allTests {
		t.Run(tc.name, func(t *testing.T) {
			have := trebuchet.DecodeAll(tc.input)
			if tc.want != have {
				t.Errorf("Decoding all failed. Want %v, have %v", tc.want, have)
			}
		})
	}
}

func TestDecodeWords(t *testing.T) {
	for _, tc := range extraCases {
		t.Run(tc.name, func(t *testing.T) {
			have := trebuchet.DecodeWithWords(tc.input)
			if tc.want != have {
				t.Errorf("Decoding failed. Want %v, have %v", tc.want, have)
			}
		})
	}
}

func TestDecodeAllWords(t *testing.T) {
	//for _, tc := range allWords {
	t.Run(allWords.name, func(t *testing.T) {
		have := trebuchet.DecodeAllWords(allWords.input)
		if allWords.want != have {
			t.Errorf("Decoding failed. Want %v, have %v", allWords.want, have)
		}
	})
	//}
}

func BenchmarkDecodeAllWords(b *testing.B) {
	/*for _, tc := range allWords {
		b.Run(tc.name, func(b *testing.B) {
			_ = trebuchet.DecodeAllWords(tc.input)
		})
	}*/
	for i := 0; i < b.N; i++ {
		_ = trebuchet.DecodeAllWords(allWords.input)
	}
}

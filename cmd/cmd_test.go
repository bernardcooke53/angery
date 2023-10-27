package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type transformTextTestCase struct {
	input    string
	expected string
}

var transformTextTestCases = []transformTextTestCase{
	{"abcd", "a b c d"},
	{"abcdefghijk word", "a b c d e f g h i j k   w o r d"},
	{"punctuation's very tricky", "p u n c t u a t i o n ' s   v e r y   t r i c k y"},
}

func TestTransformText(t *testing.T) {
	for nr, testCase := range transformTextTestCases {
		l := transformText(testCase.input)
		assert.Equal(t, testCase.expected, l, "Case %d failed", nr)
	}
}

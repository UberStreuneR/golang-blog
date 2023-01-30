package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func increment(num int) int {
	return num + 1
}

type TestNum struct {
	num      int
	expected int
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)
	cases := []TestNum{
		{num: 1, expected: 2},
		{num: 2, expected: 3},
		{num: 3, expected: 4},
		{num: 4, expected: 5},
	}
	for _, tc := range cases {
		actual := increment(tc.num)
		assert.Equal(tc.expected, actual)
	}
}

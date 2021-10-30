package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTestCase struct {
	x        int
	y        int
	expected int
}

func TestAdd(t *testing.T) {
	for _, tc := range []addTestCase{
		addTestCase{1, 2, 3},
		addTestCase{2, 2, 4},
	} {
		t.Run("add", func(t *testing.T) {
			assert.Equal(t, tc.expected, Add(tc.x, tc.y))
		})
	}
}

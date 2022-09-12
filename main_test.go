package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_display(t *testing.T) {

	n := randomNumbers()

	assert.GreaterOrEqual(t, n, 0)

	assert.LessOrEqual(t, n, 1000)
}

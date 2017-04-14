package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPermission(t *testing.T) {
	case1 := HasPermission(users[0], resources[0], operations[0])
	assert.Equal(t, case1, true, "case1 should be true")
	case2 := HasPermission(users[0], resources[0], operations[2])
	assert.Equal(t, case2, false, "case2 should be false")

	case3 := HasPermission(users[1], resources[0], operations[0])
	assert.Equal(t, case3, true, "case1 should be true")
}

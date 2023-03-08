package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	a := 4
	b := 5
	expectedResult := a + b
	require.Equal(t, add(a, b), expectedResult)
}

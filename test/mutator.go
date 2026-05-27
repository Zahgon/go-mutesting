package test

import (
	"testing"

	"github.com/zimmski/go-mutesting/mutator"
)

// Mutator tests a mutator.
// It mutates the given original file with the given mutator. Every mutation is then validated with the given changed file. The mutation overall count is validated with the given count.
func Mutator(t *testing.T, m mutator.Mutator, testFile string, count int) {
	_ = "STUB: not implemented"
	// Test if mutator is not nil
	return
}

// Read the origianl source code

// Parse and type-check the original source code

// Mutate a non relevant node

// Count the actual mutations

// Mutate all relevant nodes -> test whole mutation process

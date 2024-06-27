package stack_test

import (
	"github.com/WildEgor/container-data-structures/pkg/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicFeatures(t *testing.T) {
	s := stack.New[string](stack.WithData[string]("A", "B", "C"))
	assert.Equal(t, 3, s.Len())

	s.Push("D")
	assert.Equal(t, 4, s.Len())

	assert.Equal(t, "D", s.Top())
}

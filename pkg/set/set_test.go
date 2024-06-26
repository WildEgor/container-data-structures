package set_test

import (
	"github.com/WildEgor/container-data-structures/pkg/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicFeatures(t *testing.T) {
	s := set.New[string]("A")

	s.Add("A")

	assert.Equal(t, true, s.Has("A"))
	assert.Equal(t, s.Len(), 1)

	s.Delete("A")

	assert.Equal(t, false, s.Contains("A"))
	assert.Equal(t, s.Len(), 0)
}

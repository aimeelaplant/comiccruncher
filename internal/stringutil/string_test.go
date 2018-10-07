package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasAnyiPrefix(t *testing.T) {
	s := "jEan Grey (sjdsfjsd)"
	prefix := "Jean Grey"
	assert.True(t, HasAnyiPrefix(s, prefix))
	s1 := "No Storm"
	prefix = "Storm"
	assert.False(t, HasAnyiPrefix(s1, prefix))
}

func TestRandString(t *testing.T) {
	s := RandString(26)
	assert.NotEmpty(t, s)
	assert.Len(t, s, 26)

	s = RandString(32)
	assert.NotEmpty(t, s)
	assert.Len(t, s, 32)
}

func TestEqualsIAny(t *testing.T) {
	assert.True(t, EqualsIAny("test string ", "another", " test string"))
	assert.False(t, EqualsIAny("no ", "yess", " yesno"))
	assert.True(t, EqualsIAny("dc ", "dc comics", " dc"))
}

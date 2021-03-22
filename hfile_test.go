package hfile_test

import (
	"testing"

	"github.com/masterZSH/hfile"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash, err := hfile.Hash("/data/test/test")
	assert.Nil(t, err)
	assert.Empty(t, hash)
}

package hfile_test

import (
	"testing"

	"github.com/masterZSH/hfile"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	// invalid file
	_, err := hfile.Hash("invalid file")
	assert.NotNil(t, err)

	hash, err := hfile.Hash("/data/test/test")
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

	hashStr1, _ := hfile.HashString("/data/test/test")
	hashStr2, _ := hfile.HashString("/data/test/test")
	assert.Equal(t, hashStr1, hashStr2)
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hfile.Hash("/data/test/test")
	}
}

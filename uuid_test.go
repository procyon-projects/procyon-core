package core

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	var zero [36]byte
	var uuid [36]byte
	GenerateUUID(uuid[:])
	assert.True(t, bytes.Compare(uuid[:], zero[:]) != 0)
}

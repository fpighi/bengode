package encoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodingToString(t *testing.T) {
	bencodedString, err := Marshal("pippo")

	assert.Equal(t, "5:pippo", bencodedString)
	assert.Nil(t, err)
}

func TestEncodingEmptyString(t *testing.T) {
	bencodeString, err := Marshal("")

	assert.Equal(t, "", bencodeString)
	assert.Nil(t, err)
}

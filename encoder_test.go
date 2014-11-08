package encoder

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodingToString(t *testing.T) {
	bencodedString, err := Marshal("pippo")

	assert.Equal(t, "5:pippo", bencodedString)
	assert.Nil(t, err)
}

func TestEncodingEmptyString(t *testing.T) {
	bencodedString, err := Marshal("")

	assert.Equal(t, "", bencodedString)
	assert.Nil(t, err)
}

func TestEncodingInt8(t *testing.T) {
	v := int8(123)
	bencodedInt, err := Marshal(v)

	assert.Equal(t, "i123e", bencodedInt)
	assert.Nil(t, err)
}

func TestEncodingInt16(t *testing.T) {
	v := int16(123)
	bencodedInt, err := Marshal(v)

	assert.Equal(t, "i123e", bencodedInt)
	assert.Nil(t, err)
}

func TestEncodingInt32(t *testing.T) {
	v := int32(123)
	bencodedInt, err := Marshal(v)

	assert.Equal(t, "i123e", bencodedInt)
	assert.Nil(t, err)
}

func TestEncodingInt64(t *testing.T) {
	v := int64(123)
	bencodedInt, err := Marshal(v)

	assert.Equal(t, "i123e", bencodedInt)
	assert.Nil(t, err)
}

func TestEncodingNegativeNumber(t *testing.T) {
	v := -42
	bencodedInt, err := Marshal(v)

	assert.Equal(t, "i-42e", bencodedInt)
	assert.Nil(t, err)
}

func TestEncodingList(t *testing.T) {
	l := list.New()
	l.PushBack("pippo")
	l.PushBack(123)

	bencodedList, err := Marshal(l)

	assert.Equal(t, "l5:pippoi123ee", bencodedList)
	assert.Nil(t, err)
}

func BenchmarkEncodingToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Marshal("pippo")
	}
}

package encoder

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEncodingToString(t *testing.T) {
	bencodedString, err := Marshal("pippo")

	assert.Nil(t, err)
	assert.Equal(t, "5:pippo", bencodedString)
}

func TestEncodingEmptyString(t *testing.T) {
	bencodedString, err := Marshal("")

	assert.Nil(t, err)
	assert.Equal(t, "", bencodedString)
}

func TestEncodingInt8(t *testing.T) {
	v := int8(123)
	bencodedInt, err := Marshal(v)

	assert.Nil(t, err)
	assert.Equal(t, "i123e", bencodedInt)
}

func TestEncodingInt16(t *testing.T) {
	v := int16(123)
	bencodedInt, err := Marshal(v)

	assert.Nil(t, err)
	assert.Equal(t, "i123e", bencodedInt)
}

func TestEncodingInt32(t *testing.T) {
	v := int32(123)
	bencodedInt, err := Marshal(v)

	assert.Nil(t, err)
	assert.Equal(t, "i123e", bencodedInt)
}

func TestEncodingInt64(t *testing.T) {
	v := int64(123)
	bencodedInt, err := Marshal(v)

	assert.Nil(t, err)
	assert.Equal(t, "i123e", bencodedInt)
}

func TestEncodingNegativeNumber(t *testing.T) {
	v := -42
	bencodedInt, err := Marshal(v)

	assert.Nil(t, err)
	assert.Equal(t, "i-42e", bencodedInt)
}

func TestEncodingList(t *testing.T) {
	l := list.New()
	l.PushBack("pippo")
	l.PushBack(123)

	bencodedList, err := Marshal(l)

	assert.Nil(t, err)
	assert.Equal(t, "l5:pippoi123ee", bencodedList)
}

func TestEncodingDictionaryOfInt(t *testing.T) {
	x := map[string]interface{}{
		"foo": 1,
		"bar": 2,
	}

	bencodedDictionary, err := Marshal(x)

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(bencodedDictionary, "d"))
	assert.True(t, strings.HasSuffix(bencodedDictionary, "e"))
	assert.Contains(t, bencodedDictionary, "3:fooi1e")
	assert.Contains(t, bencodedDictionary, "3:bari2e")
	assert.Len(t, bencodedDictionary, 18)
}

func BenchmarkEncodingToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Marshal("pippo")
	}
}

package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuyaprgrm/text2speech/pkg/collection"
)

func TestSetAdd(t *testing.T) {
	set := collection.NewSet[int32]()
	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.True(t, set.Contains(2))
}

func TestRemove(t *testing.T) {
	set := collection.NewSet[int32]()
	set.Add(1)
	set.Add(2)
	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.True(t, set.Contains(2))
}

func TestLen(t *testing.T) {
	set := collection.NewSet[int32]()
	assert.Equal(t, 0, set.Len())
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(2)
	assert.Equal(t, 2, set.Len())
	set.Remove(1)
	assert.Equal(t, 1, set.Len())
}

func TestToSlice(t *testing.T) {
	set := collection.NewSet[int32]()
	set.Add(1)
	set.Add(2)
	slice := set.ToSlice()
	assert.ElementsMatch(t, []int32{1, 2}, slice)
}

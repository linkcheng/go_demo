package mydata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -run '' -v

func TestTList(t *testing.T) {
	t.Run("aaaa", func(t *testing.T) {
		l := CreateList()
		l.Append(0)
		l.Append(1)
		l.Append(2)
		l.Append(3)
		assert.Equal(t, "0,1,2,3", l.ToString())
	})

	t.Run("iiii", func(t *testing.T) {
		l := CreateList()
		l.Insert(0, 0)
		l.Insert(1, 1)
		l.Insert(2, 2)
		l.Insert(3, 3)
		assert.Equal(t, "0,1,2,3", l.ToString())
	})

	t.Run("iaia", func(t *testing.T) {
		l := CreateList()
		l.Insert(0, 0)
		l.Append(1)
		l.Insert(2, 2)
		l.Append(3)
		assert.Equal(t, "0,1,2,3", l.ToString())
	})

	t.Run("aaiid", func(t *testing.T) {
		l := CreateList()
		l.Append(1)
		l.Add(0)
		l.Insert(2, 2)
		l.Insert(3, 3)
		l.Delete(3)
		assert.Equal(t, "0,1,2", l.ToString())

		l.Reverse()
		assert.Equal(t, "2,1,0", l.ToString())
	})

	t.Run("aaiid", func(t *testing.T) {
		l := CreateList()
		l.Append(1)
		l.Add(0)
		l.Insert(2, 2)
		l.Insert(3, 3)
		l.Delete(2)
		assert.Equal(t, "0,1,3", l.ToString())
	})
}

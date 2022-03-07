package mydata

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	t.Run("aaaa", func(t *testing.T) {
		s := NewSkipList()
		
		s.Insert(10.0, 10)
		s.Insert(20.0, 20)
		s.Insert(5.0, 5)
		s.Insert(30.0, 30)
		s.Insert(15.0, 15)
		s.Insert(25.0, 25)
		println()
		s.Scan()

		e, b := s.Search(5.0)
		fmt.Printf("insert e.value=%v, exist=%t\n", e.Value, b)

		e, b = s.Search(35.0)
		fmt.Printf("insert e.value=%+v, exist=%t\n", e, b)

		e = s.Delete(5.0)
		fmt.Printf("delete e.value=%v\n", e.Value)

		e = s.Delete(35.0)
		fmt.Printf("delete e.value=%+v\n", e)

		println()
		s.Scan()
	})
}
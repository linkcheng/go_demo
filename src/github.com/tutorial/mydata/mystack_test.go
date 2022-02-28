package mydata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -run='mystack_test' -bench="Benchmark_Push" -benchmem -v
// go test -run=文件名字 -bench=bench名字 -cpuprofile=生产的cprofile文件名称 文件夹

var stack *Stack = CreateStack()

func Benchmark_Push(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
}

func Benchmark_Pop(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Push("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		stack.Pop()
	}
}

func Test_ChechParenthese(t *testing.T) {
	tests := []struct{
		in string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, ts := range tests {
		actual := ChechParenthese(ts.in)
		assert.Equal(t, ts.expected, actual)
	}

	m := make(map[int][]int)
	m[1] = []int{1,2,3}
	fmt.Println(m[1])
}

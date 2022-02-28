package main

import (
	"context"
	"strings"
)

type Node interface {
	Address() string
}

type Filter func(context.Context, []Node) []Node

type SelectOptions struct {
	Filters []Filter
}

type SelectOption func(*SelectOptions) 

func WithFilter(fs ...Filter) SelectOption {
	return func(opts *SelectOptions) {
		opts.Filters = fs
	}
}


func main() {
	size := 5
	items := []string{"a", "b"}
	values := make([]string, 0, size)
	for _, item := range items {
		values = append(values, item)
	}
	for _, v := range values {
		println(v)
	}

	value := strings.Join(values, "|")
	println(value)
}
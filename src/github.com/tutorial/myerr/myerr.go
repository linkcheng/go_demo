package main

import (
	"fmt"

	xerror "github.com/pkg/errors"
)

func main() {
	fmt.Printf("err: %+v", c())
}

func a() error {
	return xerror.Wrap(fmt.Errorf("xxx"), "test")
}

func b() error {
	return a()
}

func c() error {
	return b()
}
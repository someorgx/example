package main

import (
	"fmt"

	"github.com/someorgx/foobar"
	"github.com/someorgx/thursday/hello"
)

func main() {
	fmt.Println(hello.Hello("world"))
	fmt.Println(foobar.FooBar())
}

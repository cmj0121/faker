package main

import (
	"fmt"

	"github.com/cmj0121/faker"
)

type Simple struct {
	Ignore []byte `-`
	Count  int
	Data   []byte `fake_size:"8"`
	Name   string `fake_enum:"name"`
	Domain string `fake_enum:"domain"`
}

func main() {
	simple := Simple{}
	faker.Fake(&simple)
	fmt.Printf("%#v\n", simple)
}

package main

import (
	"fmt"

	"github.com/cmj0121/faker"
)

type Simple struct {
	Ignore []byte `-`
	Count  int
	Data   []byte `fake_size:"8"`
}

func main() {
	simple := Simple{}
	faker.Fake(&simple)
	fmt.Printf("%#v\n", simple)
}

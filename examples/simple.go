package main

import (
	"fmt"

	"github.com/cmj0121/faker"
)

type Simple struct {
	Ignore []byte `-`
	Count  int
	Data   []byte `fake_size:"8"`
	Name   string `fake:"name"`
	Domain string `fake:"domain"`
	Email  string `fake:"email"`
	Lower string `fake:"lower" fake_size:"4"`
	Upper string `fake:"upper" fake_size:"12"`
	Digit string `fake:"digit" fake_size:"2"`
}

func main() {
	simple := Simple{}
	faker.MustFake(&simple)
	fmt.Printf("%#v\n", simple)
}

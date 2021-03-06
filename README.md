# Faker #
![GitHub Action](https://github.com/cmj0121/faker/workflows/ci/badge.svg)
The structure-based faker data generator.

The following is the sample code when you need to generate the fake data: import the library, run `faker.Fake` and
you will get the fake data.

```go
package main

import (
	"fmt"

	"github.com/cmj0121/faker"
)

type Simple struct {
	Ignore []byte `-`
	Count  int
	Data   string
}

func main() {
	simple := Simple{}
	faker.Fake(&simple)
	fmt.Printf("%#v\n", simple)
}
```

## Random ##
The random interface is provide the random data with specified `Seed`. By-default the faker will
call the Seed twice when load a new Random which first-time pass the time.Now().UnixNano() and
then pass the result for the Int63(). Also, you can use the slower but more secure generator, like `CryptoRandom`
to generate the random data, or define your random generator and call `SetGenerator` to replace the generator.

## Tag ##
You can define the pre-define tag in your struct which limited the data generated from the Faker

| key       | value  | description                              |
|-----------|--------|------------------------------------------|
| -         |        | ignore the public field in the structure |
| fake_size | INT    | the limited size of the fake data        |
| fake      | name   | the limited fake name                    |
|           | domain | the limited fake top domain              |
|           | email  | the random format fake email address     |
|           | lower  | choices the lower-case [a-z]             |
|           | upper  | choices the upper-case [A-Z]             |
|           | digit  | choices the digits [0-9]                 |

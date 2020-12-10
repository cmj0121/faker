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
then pass the result for the Int63()

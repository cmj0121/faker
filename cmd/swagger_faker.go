package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cmj0121/argparse"
	"github.com/cmj0121/faker"
	"gopkg.in/yaml.v3"
)

const (
	YAML = "yaml"
	JSON = "json"
)

// the faker command-line tool
type Faker struct {
	argparse.Model

	SVersion int    `short:"V" choices:"2 3" help:"swagger v2/v3"`
	Swagger  string `short:"s" help:"external swagger file"`
	Format   string `short:"f" choices:"yaml json" help:"swagger format"`

	SwaggerV2 `-`
	SwaggerV3 `-`
}

func (faker *Faker) Run() (err error) {
	parse := argparse.MustNew(faker)
	if err = parse.Run(); err == nil {
		// run the faker after command-line parsed
		if faker.Swagger != "" {
			var data []byte

			if data, err = ioutil.ReadFile(faker.Swagger); err != nil {
				err = fmt.Errorf("open swagger '%s': %v", faker.Swagger, err)
				return
			}

			var swagger interface{}
			switch faker.SVersion {
			case 2:
				swagger = &faker.SwaggerV2
			case 3:
				swagger = &faker.SwaggerV3
			default:
				err = fmt.Errorf("not support swagger version: %v", faker.SVersion)
				return
			}
			switch faker.Format {
			case YAML:
				if err = yaml.Unmarshal(data, &swagger); err != nil {
					err = fmt.Errorf("unmarshal yaml: %s", err)
					return
				}
			case JSON:
				if err = json.Unmarshal(data, &swagger); err != nil {
					err = fmt.Errorf("unmarshal yaml: %s", err)
					return
				}
			default:
				err = fmt.Errorf("not support format: %v", faker.Format)
				return
			}

			fmt.Printf("%#v\n", swagger)
		}
	}

	return
}

// the swagger data (v2)
type SwaggerV2 struct {
	Swagger string
	Info    struct {
		Title       string
		Description string
		Version     string
	}

	Host     string
	BasePath string
	Schemes  []string
}

// the swagger data (v3)
type SwaggerV3 struct {
	OpenAPI string
	Info    struct {
		Title       string
		Description string
		Version     string
	}

	Servers []struct {
		URL         string
		Description string
	}
}

func Version(parser *argparse.ArgParse) (exit bool) {
	fmt.Printf("%s\n", faker.Version())
	return
}

func main() {
	argparse.RegisterCallback(argparse.FN_VERSION, Version)

	faker := Faker{
		SVersion: 3,
		Format:   "yaml",
	}

	if err := faker.Run(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/justuswilhelm/babble"
	"log"
	"math/rand"
	"os"
	"reflect"
)

// Config stores flags
type Config struct {
	InPlace bool
	files   []string
}

var (
	babbler = babble.NewBabbler()
	config  = Config{}
)

func init() {
	flag.BoolVar(&config.InPlace, "i", false, "Perform in-place")
	flag.Parse()
	config.files = flag.Args()
}

func anonymize(value interface{}) (interface{}, error) {
	switch val := (value).(type) {
	case nil:
		return nil, nil
	case bool:
		return val, nil
	case float64:
		return rand.Float64(), nil
	case string:
		return babbler.Babble(), nil
	case map[string]interface{}:
		newObj := make(map[string]interface{})
		for k, v := range val {
			result, err := anonymize(v)
			if err != nil {
				return nil, err
			}
			newObj[k] = result
		}
		return newObj, nil
	case []interface{}:
		newObj := make([]interface{}, len(val))
		for i, v := range val {
			result, err := anonymize(v)
			if err != nil {
				return nil, err
			}
			newObj[i] = result
		}
		return newObj, nil
	default:
		return nil, fmt.Errorf("Unknown datatype in value %+v. It has the type %v",
			value, reflect.TypeOf(value))
	}
}

func outputStdout(path string, value *interface{}) error {
	log.Printf("=== %s ===\n", path)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	if err := e.Encode(value); err != nil {
		return err
	}
	log.Printf("\n")
	return nil
}

func outputInplace(path string, value *interface{}) error {
	fd, err := os.Create(path)
	defer fd.Close()

	if err != nil {
		return err
	}
	e := json.NewEncoder(fd)
	if err := e.Encode(value); err != nil {
		return err
	}
	return nil
}

func convert(path string) error {
	var val interface{}
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	d := json.NewDecoder(fd)
	if err := d.Decode(&val); err != nil {
		return err
	}
	val, err = anonymize(val)
	if err != nil {
		return err
	}

	if config.InPlace {
		outputInplace(path, &val)
	} else {
		outputStdout(path, &val)
	}
	return nil
}

func main() {
	for _, path := range config.files {
		err := convert(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

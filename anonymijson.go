package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func anonymize(value interface{}) (interface{}, error) {
	switch val := (value).(type) {
	case float64:
		return rand.ExpFloat64(), nil
	case string:
		return randStringRunes(16), nil
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

func main() {
	var val interface{}
	d := json.NewDecoder(os.Stdin)
	if err := d.Decode(&val); err != nil {
		log.Fatal(err)
	}
	val, err := anonymize(val)
	if err != nil {
		log.Fatal(err)
	}
	e := json.NewEncoder(os.Stdout)
	if err := e.Encode(val); err != nil {
		log.Fatal(err)
	}
}

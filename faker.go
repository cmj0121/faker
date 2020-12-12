package faker

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	// global lock for the faker which may pass the same instance
	fake_lock = sync.Mutex{}
)

// run must- prefix and raise panic when error happened
func MustFake(in interface{}) {
	if err := Fake(in); err != nil {
		panic(err)
	}
}

// fake the input instance, should pass the reference
func Fake(in interface{}) (err error) {
	fake_lock.Lock()
	defer fake_lock.Unlock()

	// the reflect.Value of the instance
	value := reflect.ValueOf(in)

	// simple check the pass instance
	switch {
	case value.Kind() != reflect.Ptr:
		err = fmt.Errorf("should pass the pointer: %T", in)
		return
	case value.Kind() == reflect.Ptr && !value.Elem().CanSet():
		err = fmt.Errorf("pass %T cannot be set", in)
		return
	}

	err = fake(value.Elem(), 0)
	return
}

func fake(value reflect.Value, size int) (err error) {
	switch kind := value.Kind(); kind {
	case reflect.Bool:
		// set the random boolean value
		value.SetBool(generator.Int63()%2 == 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// set the random int64 (may truncated)
		value.SetInt(generator.Int63())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// set the random uint64 (may truncated)
		value.SetUint(uint64(generator.Int63()))
	case reflect.Float64, reflect.Float32:
		// set the random float64 (may truncated)
		value.SetFloat(generator.Float64())
	case reflect.Complex64, reflect.Complex128:
		// set random complex128 (may truncated)
		c := complex(generator.Float64(), generator.Float64())
		value.SetComplex(c)
	case reflect.Array:
		for idx := 0; idx < value.Cap(); idx++ {
			if err = fake(value.Index(idx), 0); err != nil {
				err = fmt.Errorf("cannot set #%d on %v: %v", idx, value, err)
				return
			}
		}
	case reflect.Slice:
		length := int(generator.Int63() % FAKE_MAX_SLICE_LEN)
		if size > 0 {
			// override the length
			length = size
		}

		for idx := 0; idx < length; idx++ {
			switch {
			case idx < value.Len():
				if err = fake(value.Index(idx), 0); err != nil {
					err = fmt.Errorf("cannot set #%d on %v: %v", idx, value, err)
					return
				}
			default:
				val := reflect.New(value.Type().Elem())
				if err = fake(val.Elem(), 0); err != nil {
					err = fmt.Errorf("cannot set new instance %v: %v", val.Type(), err)
					return
				}
				value.Set(reflect.Append(value, val.Elem()))
			}
		}
	case reflect.String:
		length := int(generator.Int63() % FAKE_MAX_SLICE_LEN)
		if size > 0 {
			// override the length
			length = size
		}
		str := make([]byte, length)

		for idx := 0; idx < length; idx++ {
			// save the data to the string
			str[idx] = byte(generator.Int63())
		}
		value.SetString(string(str))
	case reflect.Struct:
		for idx := 0; idx < value.NumField(); idx++ {
			if field := value.Field(idx); field.IsValid() && field.CanSet() {
				// the field now is valid and can set
				tags := value.Type().Field(idx).Tag
				size = 0

				if strings.TrimSpace(string(tags)) == FAKE_TAG_IGNORE {
					// ignore the tag
					continue
				}

				if s := tags.Get(FAKE_TAG_SLICE_SIZE); s != "" {
					if size, err = strconv.Atoi(s); err != nil {
						// cannot convert to int
						err = fmt.Errorf("invalid %v: %v", s, err)
						return
					}
				}

				// set by each field
				if err = fake(field, size); err != nil {
					// cannot set field on structure
					return
				}
			}
		}
	default:
		err = fmt.Errorf("cannot set fake for reflect.Kind: %v", kind)
		return
	}

	return
}

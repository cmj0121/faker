package faker

import (
	"fmt"
	"reflect"
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

	err = fake(value.Elem())
	return
}

func fake(value reflect.Value) (err error) {
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
	default:
		err = fmt.Errorf("cannot set fake for reflect.Kind: %v", kind)
		return
	}

	return
}

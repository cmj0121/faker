package faker

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	fake_lock = sync.Mutex{}
)

func Fake(in interface{}) (err error) {
	value := reflect.ValueOf(in)

	if value.Kind() != reflect.Ptr {
		err = fmt.Errorf("should pass the pointer: %T", in)
		return
	} else if !value.Elem().CanSet() {
		err = fmt.Errorf("pass %T cannot be set", in)
		return
	}

	err = fake(value.Elem())
	return
}

func fake(value reflect.Value) (err error) {
	fake_lock.Lock()
	defer fake_lock.Unlock()

	switch kind := value.Kind(); kind {
	case reflect.Int, reflect.Int8:
		value.SetInt(generator.Int63())
	case reflect.Float64, reflect.Float32:
		value.SetFloat(generator.Float64())
	default:
		err = fmt.Errorf("cannot set fake for reflect.Kind: %v", kind)
		return
	}

	return
}

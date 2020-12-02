package faker

import (
	"fmt"
	"reflect"
	"sync"
)

const (
	FAKE_MAX_CAP = 32
)

var (
	fake_lock = sync.Mutex{}
)

func Fake(in interface{}) (err error) {
	fake_lock.Lock()
	defer fake_lock.Unlock()

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
	switch kind := value.Kind(); kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(generator.Int63())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(uint64(generator.Int63()))
	case reflect.Float64, reflect.Float32:
		value.SetFloat(generator.Float64())
	case reflect.String:
		size := int(generator.Int63() % FAKE_MAX_CAP)
		data := fakeBytes(size, nil)
		value.SetString(string(data))
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < value.Len(); idx++ {
			if err = fake(value.Index(idx)); err != nil {
				err = fmt.Errorf("cannot set #%d on %v: %v", idx, value.Type(), err)
				return
			}
		}
		return
	default:
		err = fmt.Errorf("cannot set fake for reflect.Kind: %v", kind)
		return
	}

	return
}

func fakeBytes(size int, pool []byte) (out []byte) {
	switch l := int64(len(pool)); l {
	case 0:
		for idx := 0; idx < size; idx++ {
			out = append(out, byte(generator.Int63()))
		}
	default:
		for idx := 0; idx < size; idx++ {
			out = append(out, pool[generator.Int63()%l])
		}
	}

	return
}

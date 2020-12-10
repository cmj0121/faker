package faker

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var (
	// global lock for the faker which may pass the same instance
	fake_lock = sync.Mutex{}
)

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
		value.SetBool(generator.Int63()%2 == 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(generator.Int63())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(uint64(generator.Int63()))
	case reflect.Float64, reflect.Float32:
		value.SetFloat(generator.Float64())
	case reflect.String:
		size := int(generator.Int63() % FAKE_MAX_SIZE)
		data := fakeBytes(size, nil)
		value.SetString(string(data))
	case reflect.Slice, reflect.Array:
		size := value.Len()
		if size == 0 {
			// override the len to FAKE_MAX_SIZE
			size = FAKE_MAX_SIZE
		}

		for idx := 0; idx < size; idx++ {
			switch {
			case idx < value.Len():
				if err = fake(value.Index(idx)); err != nil {
					err = fmt.Errorf("cannot set #%d on %v: %v", idx, value.Type(), err)
					return
				}
			default:
				val := reflect.New(value.Type().Elem())
				if err = fake(val.Elem()); err != nil {
					err = fmt.Errorf("cannot set new instance %v: %v", val.Type(), err)
					return
				}
				value.Set(reflect.Append(value, val.Elem()))
			}
		}
		return
	case reflect.Struct:
		for idx := 0; idx < value.NumField(); idx++ {
			field := value.Field(idx)

			tags := value.Type().Field(idx).Tag
			if field.IsValid() && field.CanSet() {
				switch {
				case strings.TrimSpace(string(tags)) == FAKE_IGNORE:
				case tags.Get(FAKE_KEY) == FAKE_VALUE_NAME && field.Kind() == reflect.String:
					name := string(fakeBytes(FAKE_MAX_SIZE, FAKE_NAME_POOL))
					strings.ToTitle(name)
					field.SetString(name)
				case tags.Get(FAKE_KEY) == FAKE_VALUE_EMAIL && field.Kind() == reflect.String:
					email := string(fakeBytes(FAKE_MAX_SIZE, FAKE_EMAIL_POOL))
					email += "@" + string(fakeBytes(3, FAKE_NAME_POOL)) + "."
					email += FAKE_DOMAIN_POOL[int(generator.Int63())%len(FAKE_DOMAIN_POOL)]
					field.SetString(email)
				default:
					fake(field)
				}
			}
		}
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

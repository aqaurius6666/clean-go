package teststuff

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {

	var a *int
	va := reflect.ValueOf(&a).Elem()

	v := reflect.New(va.Type().Elem())
	va.Set(v)
	t.Log(*a)
}

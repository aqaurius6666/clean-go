package teststuff

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestReflect(t *testing.T) {

	var a *int
	va := reflect.ValueOf(&a).Elem()

	v := reflect.New(va.Type().Elem())
	va.Set(v)
	t.Log(*a)
}

func TestValidate(t *testing.T) {
	type MyStruct struct {
		Field string `validate:"required"`
	}
	var a MyStruct
	a.Field = ""
	v := validator.New()
	if err := v.Struct(a); err != nil {
		t.Log(err)
	}
}

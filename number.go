package gott

import (
	"errors"
	"fmt"
)


type From struct {
	value     float64
	valueStr  string
	isNumeric bool
	err       error
}

func Number(n interface{}) *From {
	f := new(From)
	f.isNumeric = true
	switch v := n.(type) {
	case int:
		f.value = float64(v)
	case int8:
		f.value = float64(v)
	case int16:
		f.value = float64(v)
	case int32:
		f.value = float64(v)
	case int64:
		f.value = float64(v)
	case float32:
		f.value = float64(v)
	case float64:
		f.value = v
	case string:
		f.isNumeric = false
		f.valueStr = v
	default:
		f.err = errors.New(fmt.Sprintf("invalid value : %v", n))
	}
	return f
}

func (f *From) ToInt32() (int32, error) {
	return 0, nil
}

func (f *From) ToFloat32() (float32, error) {
	return 0, nil
}

//Converts number to specific culture
func (f *From) ToString(culture string) (string, error) {
	return "", nil
}

//Converts money to specific culture
func (f *From) ToMoney(culture string) (string, error) {
	if f.err != nil {
		return "", f.err
	}
	ci, ok := cultures[culture]
	if !ok {
		return "", errors.New(fmt.Sprintf("culture '%s' not supported", culture))
	}
	v := f.build(ci.NumberFormatInfo)

	return v, nil
}

//Converts number to specific culture
func (f *From) GetCultureInfo() (int, error) {
	return 0, nil
}

func parseNumber(value string, culture string) *From {
	f := new(From)
	return f
}

func (f *From) build(numberFormat numberFormatInfo) string {

	result := newBuilder(f.value, numberFormat).forehead().decimal().currency()

	return result.result
}

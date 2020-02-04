package gott

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)


type builder struct {
	numberFormat numberFormatInfo
	number float64
	result string
}


func newBuilder(number float64, numberFormat numberFormatInfo) *builder {
	builder := new(builder)
	builder.number = number
	builder.numberFormat = numberFormat

	return builder
}

func (b *builder) forehead() *builder {

	left := []byte(strconv.Itoa(int(b.number)))

	result := []byte(b.result)

	to := len(left)

	for i:= 0; true; i++ {

		grp := getGroupSize(i,b.numberFormat.CurrencyGroupSizes)

		if grp == 0 {
			result = append(left,result... )
			break
		}

		from := to - grp

		if from <= 0 {
			result = append(left[0:to],result... )
			break
		}

		result = append(left[from:to],result... )

		result = append([]byte(b.numberFormat.CurrencyGroupSeparator),result... )

		to = from
	}

	b.result = string(result)

	return b
}


func (b *builder) decimal() *builder{

	dv :=  b.number - float64(int(b.number))

	o := 1
	if b.number < 0 {
		o = 0
	}

	digits := strconv.Itoa(int( math.Round((dv+float64(o))*math.Pow10(int(b.numberFormat.CurrencyDecimalDigits)))) )

	digits= string([]byte(digits)[1:len(digits)])

	b.result = fmt.Sprintf("%s%s%v",b.result,b.numberFormat.CurrencyDecimalSeparator,digits)

	return b
}


func (b *builder) currency() *builder{
	//todo: move to newBuilder ?
	pattern := currencyPositivePattern[b.numberFormat.CurrencyPositivePattern]
	if b.number < 0 {
		pattern = currencyNegativePattern[b.numberFormat.CurrencyNegativePattern]
	}

	b.result = strings.Replace(pattern,"n", b.result,-1)

	b.result = strings.Replace(b.result,"$", b.numberFormat.CurrencySymbol,-1)



	return b
}


func getGroupSize(index int, groupSizes []int32) int {
	var grp int32
	l := len(groupSizes)
	if index < l {
		grp = groupSizes[index]
	} else {
		grp =  groupSizes[l-1]
	}
	return int(grp)
}

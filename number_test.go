package gott

import (
	"gott/codes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntAsUsCurrency(t *testing.T){

	//act
	currency200, err200 := Number(200).ToMoney(codes.EnglishUnitedStates)
	currency2001, err2001 := Number(2001).ToMoney(codes.EnglishUnitedStates)
	currency1999, err1999 := Number(1999).ToMoney(codes.EnglishUnitedStates)
	currency0, err0 := Number(0).ToMoney(codes.EnglishUnitedStates)
	currencym1, errm1 := Number(-1).ToMoney(codes.EnglishUnitedStates)


	//assert
	assert.NoError(t, err200)
	assert.NoError(t, err2001)
	assert.NoError(t, err1999)
	assert.NoError(t, err0)
	assert.NoError(t, errm1)

	assert.Equal(t, "$200.00", currency200)
	assert.Equal(t, "$2,001.00", currency2001)
	assert.Equal(t, "$1,999.00", currency1999)
	assert.Equal(t, "$0.00", currency0)
	assert.Equal(t, "($-1.00)", currencym1)
}



func TestFloatAsUsCurrency(t *testing.T){

	//act
	currency1999x15, err1999x15 := Number(1999.15).ToMoney(codes.EnglishUnitedStates)
	currency3x151, err3x151 := Number(3.151).ToMoney(codes.EnglishUnitedStates)
	currency0x0, err0x0 := Number(0.00).ToMoney(codes.EnglishUnitedStates)
	currency0x5194445, err0x5194445 := Number(0.5194445).ToMoney(codes.EnglishUnitedStates)
	currency0x5144445, err0x5144445 := Number(0.5144445).ToMoney(codes.EnglishUnitedStates)
	currencym0x166, errm0x166 := Number(-0.166).ToMoney(codes.EnglishUnitedStates)

	//assert
	assert.NoError(t, err1999x15)
	assert.NoError(t, err3x151)
	assert.NoError(t, err0x0)
	assert.NoError(t, err0x5194445)
	assert.NoError(t, err0x5144445)
	assert.NoError(t, errm0x166)

	assert.Equal(t, "$1,999.15", currency1999x15)
	assert.Equal(t, "$3.15", currency3x151)
	assert.Equal(t, "$0.00", currency0x0)
	assert.Equal(t, "$0.52", currency0x5194445)
	assert.Equal(t, "$0.51", currency0x5144445)
	assert.Equal(t, "($0.17)", currencym0x166)
}




func TestFloatAsCzCurrency(t *testing.T){

	//act
	currency1545233x156, err1545233x156 := Number(1545233.156).ToMoney(codes.CzechCzechRepublic)

	//assert
	assert.NoError(t, err1545233x156)

	assert.Equal(t, "1 545 233,16 Kč", currency1545233x156)
}




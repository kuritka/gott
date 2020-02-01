package gott

type NumberFormatInfo struct {

	//The number of decimal places to use in currency values
	//CurrencyDecimalDigits = 2: $1,234.00
	//CurrencyDecimalDigits = 4: $1,234.0000
	CurrencyDecimalDigits int32

	//The string to use as the decimal separator in currency values
	//CurrencyDecimalSeparator = ".";  $123,456,789.00
	//CurrencyDecimalSeparator = " "; $123,456,789 00
	CurrencyDecimalSeparator string

	//The string that separates groups of digits to the left of the decimal in currency values
	//CurrencyDecimalSeparator = ",";  $123,456,789.00
	//CurrencyDecimalSeparator = " "; $123 456 789.00
	CurrencyGroupSeparator string

	//The number of digits in each group to the left of the decimal in currency values
	//CurrencyGroupSizes = []{ 3, 4, 5 }; $55,55555,55555,55555,4444,333.00
	//CurrencyGroupSizes = []{ 3, 4, 0 }; $55555555555555555,4444,333.00
	CurrencyGroupSizes []int32

	//The format pattern for negative currency
	//0	($n)
	//1	-$n
	//2	$-n
	//3	$n-
	//4	(n$)
	//5	-n$
	//6	n-$
	//7	n$-
	//8	-n $
	//9	-$ n
	//10	n $-
	//11	$ n-
	//12	$ -n
	//13	n- $
	//14	($ n)
	//15	(n $)
	CurrencyNegativePattern int32

	//0	$n
	//1	n$
	//2	$ n
	//3	n $
	CurrencyPositivePattern int32

	//The string to use as the currency symbol
	//CurrencySymbol="$"; $106.25
	CurrencySymbol string

	//The string that represents negative infinity.
	//-Infinity
	NegativeInfinitySymbol string

	//The string that denotes that the associated number is negative
	NegativeSign string

	//The number of decimal places to use in numeric values
	//NumberDecimalDigits=2; -1,234.00
	//NumberDecimalDigits=4; -1,234.0000
	NumberDecimalDigits int32

	//the number decimal separator
	NumberDecimalSeparator string

	//The string that separates groups of digits to the left of the decimal in numeric values
	//NumberGroupSeparator = ","; 123,456,789.00
	//NumberGroupSeparator = " "; 123 456 789.00
	NumberGroupSeparator string

	//The number of digits in each group to the left of the decimal in numeric values
	//NumberGroupSizes = []{ 3, 4, 5 }; $55,55555,55555,55555,4444,333.00
	//NumberGroupSizes = []{ 3, 4, 0 }; $55555555555555555,4444,333.00
	NumberGroupSizes []int32

	//The format pattern for negative numeric values
	//       Default:             -1,234.00
	//       Pattern 0:           (1,234.00)
	//       Pattern 1:           -1,234.00
	//       Pattern 2:           - 1,234.00
	//       Pattern 3:           1,234.00-
	//       Pattern 4:           1,234.00 -
	NumberNegativePattern int32

	//The number of decimal places to use in percent values
	//PercentDecimalDigits = 2; 12.34 %
	//PercentDecimalDigits = 4; 12.3400 %
	PercentDecimalDigits int32

	//The string to use as the decimal separator in percent values
	//12.34 %
	//12 34 %
	PercentDecimalSeparator string

	//The string that separates groups of digits to the left of the decimal in percent values
	//PercentGroupSeparator= ",";  123,456.78 %
	//PercentGroupSeparator= " "; 123 456.78 %
	PercentGroupSeparator string

	//The number of digits in each group to the left of the decimal in percent values
	//PercentGroupSizes = { 3, 4, 5 }; 55,55555,55555,55555,4444,333.00%
	//PercentGroupSizes = { 3, 4, 0 }; 55555555555555555,4444,333.00%
	//{2,3,0}; 123456789012,346,00.00 %
	//{2,3,4}; 1234,5678,9012,346,00.00 %
	PercentGroupSizes []int32

	//0	-n %
	//1	-n%
	//2	-%n
	//3	%-n
	//4	%n-
	//5	n-%
	//6	n%-
	//7	-% n
	//8	n %-
	//9	% n-
	//10	% -n
	//11	n- %
	PercentNegativePattern int32

	//0	n %
	//1	n%
	//2	%n
	//3	% n
	PercentPositivePattern int32

	//The string to use as the percent symbol.
	PercentSymbol string

	//The string to use as the per mille symbol "‰" (U+2030)
	PerMilleSymbol string

	//The string that represents positive infinity, "Infinity"
	PositiveInfinitySymbol string

	//The string that denotes that the associated number is positive
	PositiveSign string
}

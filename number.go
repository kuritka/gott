package gott

//ToMoney
//ToString
//tt.CultureInfo() //gets culture info
//cultureInfo, DatetimeFormatInfo, NumberFormatInfo

//current culture


type FromCI struct {
	culture CultureInfo
	value interface{}

}

func ParseNumber(value string, culture string) *FromCI {
	f := new(FromCI)
	ci := cultures[culture]
	f.culture = ci

}


func Number(n interface{}) *FromCI{
	return new(FromCI)
}


func (f *FromCI) ToInt32() (int32, error) {
	return 0, nil
}


func (f *FromCI) ToFloat32() (float32, error) {
	return 0, nil
}

//Converts number to specific culture
func (f *FromCI) ToString(culture string) (string, error) {
	return "", nil
}

//Converts money to specific culture
func (f *FromCI) ToMoney(culture string) (string, error) {
	return "", nil
}

//Converts number to specific culture
func (f *FromCI) GetCultureInfo() (int, error) {
	return 0, nil
}
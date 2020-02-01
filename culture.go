package gott

type CultureInfo struct {

	//culture name
	Culture string

	//the full name of culture
	Name    string

	//culture code
	Code    int

	//culture ISO code
	ISO     string


}

var cultures = map[string]CultureInfo{
		"af-ZA":    {Culture: "af-ZA", Name: "Afrikaans - South Africa", Code: 0x0436, ISO: "AFK"},
		"sq-AL":    {Culture: "sq-AL", Name: "Albanian - Albania", Code: 0x041C, ISO: "SQI"},
		"ar-DZ":    {Culture: "ar-DZ", Name: "Arabic - Algeria", Code: 0x1401, ISO: "ARG"},
		"ar-BH":    {Culture: "ar-BH", Name: "Arabic - Bahrain", Code: 0x3C01, ISO: "ARH"},
		"ar-EG":    {Culture: "ar-EG", Name: "Arabic - Egypt", Code: 0x0C01, ISO: "ARE"},
	"ar-IQ":    {Culture: "ar-IQ", Name: "Arabic - Iraq", Code: 0x0801, ISO: "ARI"},
	"ar-JO":    {Culture: "ar-JO", Name: "Arabic - Jordan", Code: 0x2C01, ISO: "ARJ"},
	"ar-KW":    {Culture: "ar-KW", Name: "Arabic - Kuwait", Code: 0x3401, ISO: "ARK"},
	"ar-LB":    {Culture: "ar-LB", Name: "Arabic - Lebanon", Code: 0x3001, ISO: "ARB"},
	"ar-LY":    {Culture: "ar-LY", Name: "Arabic - Libya", Code: 0x1001, ISO: "ARL"},
	"ar-MA":    {Culture: "ar-MA", Name: "Arabic - Morocco", Code: 0x1801, ISO: "ARM"},
	"ar-OM":    {Culture: "ar-OM", Name: "Arabic - Oman", Code: 0x2001, ISO: "ARO"},
	"ar-QA":    {Culture: "ar-QA", Name: "Arabic - Qatar", Code: 0x4001, ISO: "ARQ"},
	"ar-SA":    {Culture: "ar-SA", Name: "Arabic - Saudi Arabia", Code: 0x0401, ISO: "ARA"},
	"ar-SY":    {Culture: "ar-SY", Name: "Arabic - Syria", Code: 0x2801, ISO: "ARS"},
	"ar-TN":    {Culture: "ar-TN", Name: "Arabic - Tunisia", Code: 0x1C01, ISO: "ART"},
	"ar-AE":    {Culture: "ar-AE", Name: "Arabic - United Arab Emirates", Code: 0x3801, ISO: "ARU"},
	"ar-YE":    {Culture: "ar-YE", Name: "Arabic - Yemen", Code: 0x2401, ISO: "ARY"},
	"hy-AM":    {Culture: "hy-AM", Name: "Armenian - Armenia", Code: 0x042B, ISO: ""},
	"Cy-az-AZ": {Culture: "Cy-az-AZ", Name: "Azeri (Cyrillic) - Azerbaijan", Code: 0x082C, ISO: ""},
	"Lt-az-AZ": {Culture: "Lt-az-AZ", Name: "Azeri (Latin) - Azerbaijan", Code: 0x042C, ISO: ""},
	"eu-ES":    {Culture: "eu-ES", Name: "Basque - Basque", Code: 0x042D, ISO: "EUQ"},
	"be-BY":    {Culture: "be-BY", Name: "Belarusian - Belarus", Code: 0x0423, ISO: "BEL"},
	"bg-BG":    {Culture: "bg-BG", Name: "Bulgarian - Bulgaria", Code: 0x0402, ISO: "BGR"},
	"ca-ES":    {Culture: "ca-ES", Name: "Catalan - Catalan", Code: 0x0403, ISO: "CAT"},
	"zh-CN":    {Culture: "zh-CN", Name: "Chinese - China", Code: 0x0804, ISO: "CHS"},
	"zh-HK":    {Culture: "zh-HK", Name: "Chinese - Hong Kong SAR", Code: 0x0C04, ISO: "ZHH"},
	"zh-MO":    {Culture: "zh-MO", Name: "Chinese - Macau SAR", Code: 0x1404, ISO: ""},
	"zh-SG":    {Culture: "zh-SG", Name: "Chinese - Singapore", Code: 0x1004, ISO: "ZHI"},
	"zh-TW":    {Culture: "zh-TW", Name: "Chinese - Taiwan", Code: 0x0404, ISO: "CHT"},
	"zh-CHS":   {Culture: "zh-CHS", Name: "Chinese (Simplified)", Code: 0x0004, ISO: ""},
	"zh-CHT":   {Culture: "zh-CHT", Name: "Chinese (Traditional)", Code: 0x7C04, ISO: ""},
	"hr-HR":    {Culture: "hr-HR", Name: "Croatian - Croatia", Code: 0x041A, ISO: "HRV"},
	"cs-CZ":    {Culture: "cs-CZ", Name: "Czech - Czech Republic", Code: 0x0405, ISO: "CSY"},
	"da-DK":    {Culture: "da-DK", Name: "Danish - Denmark", Code: 0x0406, ISO: "DAN"},
	"div-MV":   {Culture: "div-MV", Name: "Dhivehi - Maldives", Code: 0x0465, ISO: ""},
	"nl-BE":    {Culture: "nl-BE", Name: "Dutch - Belgium", Code: 0x0813, ISO: "NLB"},
	"nl-NL":    {Culture: "nl-NL", Name: "Dutch - The Netherlands", Code: 0x0413, ISO: ""},
	"en-AU":    {Culture: "en-AU", Name: "English - Australia", Code: 0x0C09, ISO: "ENA"},
	"en-BZ":    {Culture: "en-BZ", Name: "English - Belize", Code: 0x2809, ISO: "ENL"},
	"en-CA":    {Culture: "en-CA", Name: "English - Canada", Code: 0x1009, ISO: "ENC"},
	"en-CB":    {Culture: "en-CB", Name: "English - Caribbean", Code: 0x2409, ISO: "ENB"},
	"en-IE":    {Culture: "en-IE", Name: "English - Ireland", Code: 0x1809, ISO: "ENI"},
	"en-JM":    {Culture: "en-JM", Name: "English - Jamaica", Code: 0x2009, ISO: "ENJ"},
	"en-NZ":    {Culture: "en-NZ", Name: "English - New Zealand", Code: 0x1409, ISO: "ENZ"},
	"en-PH":    {Culture: "en-PH", Name: "English - Philippines", Code: 0x3409, ISO: "ENP"},
	"en-ZA":    {Culture: "en-ZA", Name: "English - South Africa", Code: 0x1C09, ISO: "ENS"},
	"en-TT":    {Culture: "en-TT", Name: "English - Trinidad and Tobago", Code: 0x2C09, ISO: "ENT"},
	"en-GB":    {Culture: "en-GB", Name: "English - United Kingdom", Code: 0x0809, ISO: "ENG"},
	"en-US":    {Culture: "en-US", Name: "English - United States", Code: 0x0409, ISO: "ENU"},
	"en-ZW":    {Culture: "en-ZW", Name: "English - Zimbabwe", Code: 0x3009, ISO: "ENW"},
	"et-EE":    {Culture: "et-EE", Name: "Estonian - Estonia", Code: 0x0425, ISO: "ETI"},
	"fo-FO":    {Culture: "fo-FO", Name: "Faroese - Faroe Islands", Code: 0x0438, ISO: "FOS"},
	"fa-IR":    {Culture: "fa-IR", Name: "Farsi - Iran", Code: 0x0429, ISO: "FAR"},
	"fi-FI":    {Culture: "fi-FI", Name: "Finnish - Finland", Code: 0x040B, ISO: "FIN"},
	"fr-BE":    {Culture: "fr-BE", Name: "French - Belgium", Code: 0x080C, ISO: "FRB"},
	"fr-CA":    {Culture: "fr-CA", Name: "French - Canada", Code: 0x0C0C, ISO: "FRC"},
	"fr-FR":    {Culture: "fr-FR", Name: "French - France", Code: 0x040C, ISO: "FRA"},
	"fr-LU":    {Culture: "fr-LU", Name: "French - Luxembourg", Code: 0x140C, ISO: "FRL"},
	"fr-MC":    {Culture: "fr-MC", Name: "French - Monaco", Code: 0x180C, ISO: "FRM"},
	"fr-CH":    {Culture: "fr-CH", Name: "French - Switzerland", Code: 0x100C, ISO: "FRS"},
	"gl-ES":    {Culture: "gl-ES", Name: "Galician - Galician", Code: 0x0456, ISO: ""},
	"ka-GE":    {Culture: "ka-GE", Name: "Georgian - Georgia", Code: 0x0437, ISO: ""},
	"de-AT":    {Culture: "de-AT", Name: "German - Austria", Code: 0x0C07, ISO: "DEA"},
	"de-DE":    {Culture: "de-DE", Name: "German - Germany", Code: 0x0407, ISO: "DEU"},
	"de-LI":    {Culture: "de-LI", Name: "German - Liechtenstein", Code: 0x1407, ISO: "DEC"},
	"de-LU":    {Culture: "de-LU", Name: "German - Luxembourg", Code: 0x1007, ISO: "DEL"},
	"de-CH":    {Culture: "de-CH", Name: "German - Switzerland", Code: 0x0807, ISO: "DES"},
	"el-GR":    {Culture: "el-GR", Name: "Greek - Greece", Code: 0x0408, ISO: "ELL"},
	"gu-IN":    {Culture: "gu-IN", Name: "Gujarati - India", Code: 0x0447, ISO: ""},
	"he-IL":    {Culture: "he-IL", Name: "Hebrew - Israel", Code: 0x040D, ISO: "HEB"},
	"hi-IN":    {Culture: "hi-IN", Name: "Hindi - India", Code: 0x0439, ISO: "HIN"},
	"hu-HU":    {Culture: "hu-HU", Name: "Hungarian - Hungary", Code: 0x040E, ISO: "HUN"},
	"is-IS":    {Culture: "is-IS", Name: "Icelandic - Iceland", Code: 0x040F, ISO: "ISL"},
	"id-ID":    {Culture: "id-ID", Name: "Indonesian - Indonesia", Code: 0x0421, ISO: "IND"},
	"default":  {Culture: "default", Name: "Default culture", Code: 0x0, ISO: ""},
}


func GetCurrentCulture() CultureInfo {
	return cultures["default"]
}

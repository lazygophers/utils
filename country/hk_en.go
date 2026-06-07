package country

import xlanguage "golang.org/x/text/language"

// English is one of Hong Kong's official languages; this file is unguarded.
func init() {
	dataHongKong.RegisterName(xlanguage.English, "Hong Kong")
	dataHongKong.RegisterOfficialName(xlanguage.English, "Hong Kong Special Administrative Region of the People's Republic of China")
	dataHongKong.RegisterCapital(xlanguage.English, "Hong Kong")
}

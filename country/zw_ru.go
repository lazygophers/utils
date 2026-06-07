//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Russian, "Зимбабве")
	dataZimbabwe.RegisterOfficialName(xlanguage.Russian, "Республика Зимбабве")
	dataZimbabwe.RegisterCapital(xlanguage.Russian, "Хараре")
}

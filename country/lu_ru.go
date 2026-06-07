//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Russian, "Люксембург")
	dataLuxembourg.RegisterOfficialName(xlanguage.Russian, "Великое Герцогство Люксембург")
	dataLuxembourg.RegisterCapital(xlanguage.Russian, "Люксембург")
}

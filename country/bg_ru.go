//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Russian, "Болгария")
	dataBulgaria.RegisterOfficialName(xlanguage.Russian, "Республика Болгария")
	dataBulgaria.RegisterCapital(xlanguage.Russian, "София")
}

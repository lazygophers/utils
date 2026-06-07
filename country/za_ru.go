//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.Russian, "Южно-Африканская Республика")
	dataSouthAfrica.RegisterOfficialName(xlanguage.Russian, "Южно-Африканская Республика")
	dataSouthAfrica.RegisterCapital(xlanguage.Russian, "Претория")
}

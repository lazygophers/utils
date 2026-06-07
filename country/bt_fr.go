//go:build (lang_fr || lang_all) && (country_all || country_asia || country_bt || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.French, "Bhoutan")
	dataBhutan.RegisterOfficialName(xlanguage.French, "Royaume du Bhoutan")
	dataBhutan.RegisterCapital(xlanguage.French, "Thimphou")
}

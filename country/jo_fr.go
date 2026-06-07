//go:build (lang_fr || lang_all) && (country_all || country_asia || country_jo || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.French, "Jordanie")
	dataJordan.RegisterOfficialName(xlanguage.French, "Royaume hachémite de Jordanie")
	dataJordan.RegisterCapital(xlanguage.French, "Amman")
}

//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.French, "Zambie")
	dataZambia.RegisterOfficialName(xlanguage.French, "République de Zambie")
	dataZambia.RegisterCapital(xlanguage.French, "Lusaka")
}

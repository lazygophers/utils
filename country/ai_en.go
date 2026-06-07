//go:build country_ai || country_all || country_americas || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.English, "Anguilla")
	dataAnguilla.RegisterOfficialName(xlanguage.English, "Anguilla")
	dataAnguilla.RegisterCapital(xlanguage.English, "The Valley")
}

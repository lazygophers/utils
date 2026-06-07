//go:build country_africa || country_all || country_dj || country_eastern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.French, "Djibouti")
	dataDjibouti.RegisterOfficialName(xlanguage.French, "République de Djibouti")
	dataDjibouti.RegisterCapital(xlanguage.French, "Djibouti")
}

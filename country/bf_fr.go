//go:build country_africa || country_all || country_bf || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.French, "Burkina Faso")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.French, "Burkina Faso")
	dataBurkinaFaso.RegisterCapital(xlanguage.French, "Ouagadougou")
}

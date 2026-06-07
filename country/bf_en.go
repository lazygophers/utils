//go:build country_africa || country_all || country_bf || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.English, "Burkina Faso")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.English, "Burkina Faso")
	dataBurkinaFaso.RegisterCapital(xlanguage.English, "Ouagadougou")
}

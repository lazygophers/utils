//go:build country_all || country_antarctic || country_aq

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.English, "Antarctica")
	dataAntarctica.RegisterOfficialName(xlanguage.English, "Antarctica")
}

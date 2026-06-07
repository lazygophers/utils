//go:build (lang_fr || lang_all) && (country_all || country_antarctic || country_aq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.French, "Antarctique")
	dataAntarctica.RegisterOfficialName(xlanguage.French, "Antarctique")
}

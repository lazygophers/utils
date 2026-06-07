//go:build (lang_fr || lang_all) && (country_all || country_europe || country_mk || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.French, "Macédoine du Nord")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.French, "République de Macédoine du Nord")
	dataNorthMacedonia.RegisterCapital(xlanguage.French, "Skopje")
}

//go:build (lang_fr || lang_all) && (country_all || country_europe || country_gr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.French, "Grèce")
	dataGreece.RegisterOfficialName(xlanguage.French, "République hellénique")
	dataGreece.RegisterCapital(xlanguage.French, "Athènes")
}

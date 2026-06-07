//go:build (lang_fr || lang_all) && (country_all || country_micronesia || country_oceania || country_pw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.French, "Palaos")
	dataPalau.RegisterOfficialName(xlanguage.French, "République des Palaos")
	dataPalau.RegisterCapital(xlanguage.French, "Ngerulmud")
}

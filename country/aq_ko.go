//go:build (lang_ko || lang_all) && (country_all || country_antarctic || country_aq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Korean, "남극")
	dataAntarctica.RegisterOfficialName(xlanguage.Korean, "남극")
}

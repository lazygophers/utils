//go:build (lang_es || lang_all) && (country_africa || country_all || country_ci || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Spanish, "Costa de Marfil")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Spanish, "República de Costa de Marfil")
	dataIvoryCoast.RegisterCapital(xlanguage.Spanish, "Yamusukro")
}

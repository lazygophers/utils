//go:build (lang_es || lang_all) && (country_africa || country_all || country_eg || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Spanish, "Egipto")
	dataEgypt.RegisterOfficialName(xlanguage.Spanish, "República Árabe de Egipto")
	dataEgypt.RegisterCapital(xlanguage.Spanish, "El Cairo")
}

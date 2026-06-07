//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eg || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.French, "Égypte")
	dataEgypt.RegisterOfficialName(xlanguage.French, "République arabe d'Égypte")
	dataEgypt.RegisterCapital(xlanguage.French, "Le Caire")
}

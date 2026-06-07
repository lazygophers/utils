//go:build country_africa || country_all || country_eg || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.English, "Egypt")
	dataEgypt.RegisterOfficialName(xlanguage.English, "Arab Republic of Egypt")
	dataEgypt.RegisterCapital(xlanguage.English, "Cairo")
}

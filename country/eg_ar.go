//go:build country_africa || country_all || country_eg || country_northern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Arabic, "مصر")
	dataEgypt.RegisterOfficialName(xlanguage.Arabic, "جمهورية مصر العربية")
	dataEgypt.RegisterCapital(xlanguage.Arabic, "القاهرة")
}

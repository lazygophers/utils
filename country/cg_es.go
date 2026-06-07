//go:build (lang_es || lang_all) && (country_africa || country_all || country_cg || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Spanish, "República del Congo")
	dataCongo.RegisterOfficialName(xlanguage.Spanish, "República del Congo")
	dataCongo.RegisterCapital(xlanguage.Spanish, "Brazzaville")
}

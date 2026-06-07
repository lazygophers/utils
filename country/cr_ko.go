//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_cr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Korean, "코스타리카")
	dataCostaRica.RegisterOfficialName(xlanguage.Korean, "코스타리카 공화국")
	dataCostaRica.RegisterCapital(xlanguage.Korean, "산호세")
}

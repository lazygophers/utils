//go:build (lang_ko || lang_all) && (country_africa || country_all || country_southern_africa || country_za)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.Korean, "남아프리카 공화국")
	dataSouthAfrica.RegisterOfficialName(xlanguage.Korean, "남아프리카 공화국")
	dataSouthAfrica.RegisterCapital(xlanguage.Korean, "프리토리아")
}

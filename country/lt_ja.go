//go:build (lang_ja || lang_all) && (country_all || country_europe || country_lt || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Japanese, "リトアニア")
	dataLithuania.RegisterOfficialName(xlanguage.Japanese, "リトアニア共和国")
	dataLithuania.RegisterCapital(xlanguage.Japanese, "ヴィリニュス")
}

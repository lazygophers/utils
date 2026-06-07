//go:build (lang_ar || lang_all) && (country_all || country_europe || country_lt || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Arabic, "ليتوانيا")
	dataLithuania.RegisterOfficialName(xlanguage.Arabic, "جمهورية ليتوانيا")
	dataLithuania.RegisterCapital(xlanguage.Arabic, "فيلنيوس")
}

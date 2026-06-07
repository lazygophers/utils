//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_km)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Japanese, "コモロ")
	dataComoros.RegisterOfficialName(xlanguage.Japanese, "コモロ連合")
	dataComoros.RegisterCapital(xlanguage.Japanese, "モロニ")
}

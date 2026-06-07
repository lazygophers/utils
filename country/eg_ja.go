//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eg || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Japanese, "エジプト")
	dataEgypt.RegisterOfficialName(xlanguage.Japanese, "エジプト・アラブ共和国")
	dataEgypt.RegisterCapital(xlanguage.Japanese, "カイロ")
}

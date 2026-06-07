//go:build (lang_ja || lang_all) && (country_africa || country_all || country_southern_africa || country_sz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Japanese, "エスワティニ")
	dataEswatini.RegisterOfficialName(xlanguage.Japanese, "エスワティニ王国")
	dataEswatini.RegisterCapital(xlanguage.Japanese, "ムババーネ")
}

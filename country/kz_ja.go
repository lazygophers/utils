//go:build (lang_ja || lang_all) && (country_all || country_asia || country_central_asia || country_kz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.Japanese, "カザフスタン")
	dataKazakhstan.RegisterOfficialName(xlanguage.Japanese, "カザフスタン共和国")
	dataKazakhstan.RegisterCapital(xlanguage.Japanese, "アスタナ")
}

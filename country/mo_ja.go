//go:build (lang_ja || lang_all) && (country_all || country_asia || country_eastern_asia || country_mo)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.Japanese, "マカオ")
	dataMacao.RegisterOfficialName(xlanguage.Japanese, "中華人民共和国マカオ特別行政区")
	dataMacao.RegisterCapital(xlanguage.Japanese, "マカオ")
}

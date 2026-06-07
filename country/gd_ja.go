//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_gd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Japanese, "グレナダ")
	dataGrenada.RegisterOfficialName(xlanguage.Japanese, "グレナダ")
	dataGrenada.RegisterCapital(xlanguage.Japanese, "セントジョージズ")
}

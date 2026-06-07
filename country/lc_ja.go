//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_lc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.Japanese, "セントルシア")
	dataSaintLucia.RegisterOfficialName(xlanguage.Japanese, "セントルシア")
	dataSaintLucia.RegisterCapital(xlanguage.Japanese, "カストリーズ")
}

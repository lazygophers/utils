//go:build (lang_ja || lang_all) && (country_ad || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.Japanese, "アンドラ")
	dataAndorra.RegisterOfficialName(xlanguage.Japanese, "アンドラ公国")
	dataAndorra.RegisterCapital(xlanguage.Japanese, "アンドラ・ラ・ベリャ")
}

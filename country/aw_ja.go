//go:build (lang_ja || lang_all) && (country_all || country_americas || country_aw || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Japanese, "アルバ")
	dataAruba.RegisterOfficialName(xlanguage.Japanese, "アルバ")
	dataAruba.RegisterCapital(xlanguage.Japanese, "オラニエスタッド")
}

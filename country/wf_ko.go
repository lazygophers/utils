//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_polynesia || country_wf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Korean, "왈리스 푸투나")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Korean, "왈리스 푸투나 준주")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Korean, "마타우투")
}

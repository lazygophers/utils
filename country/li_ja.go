//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Japanese, "リヒテンシュタイン")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Japanese, "リヒテンシュタイン公国")
	dataLiechtenstein.RegisterCapital(xlanguage.Japanese, "ファドゥーツ")
}

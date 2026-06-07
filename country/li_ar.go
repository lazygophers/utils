//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Arabic, "ليختنشتاين")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Arabic, "إمارة ليختنشتاين")
	dataLiechtenstein.RegisterCapital(xlanguage.Arabic, "فادوتس")
}

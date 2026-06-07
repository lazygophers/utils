//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Japanese, "ジョージア")
	dataGeorgia.RegisterOfficialName(xlanguage.Japanese, "ジョージア")
	dataGeorgia.RegisterCapital(xlanguage.Japanese, "トビリシ")
}

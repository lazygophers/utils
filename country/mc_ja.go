//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.Japanese, "モナコ")
	dataMonaco.RegisterOfficialName(xlanguage.Japanese, "モナコ公国")
	dataMonaco.RegisterCapital(xlanguage.Japanese, "モナコ")
}

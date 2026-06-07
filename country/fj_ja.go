//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.Japanese, "フィジー")
	dataFiji.RegisterOfficialName(xlanguage.Japanese, "フィジー共和国")
	dataFiji.RegisterCapital(xlanguage.Japanese, "スバ")
}

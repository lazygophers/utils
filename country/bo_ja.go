//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.Japanese, "ボリビア")
	dataBolivia.RegisterOfficialName(xlanguage.Japanese, "ボリビア多民族国")
	dataBolivia.RegisterCapital(xlanguage.Japanese, "スクレ")
}

//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Japanese, "ノルウェー")
	dataNorway.RegisterOfficialName(xlanguage.Japanese, "ノルウェー王国")
	dataNorway.RegisterCapital(xlanguage.Japanese, "オスロ")
}

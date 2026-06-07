//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.Japanese, "マルタ")
	dataMalta.RegisterOfficialName(xlanguage.Japanese, "マルタ共和国")
	dataMalta.RegisterCapital(xlanguage.Japanese, "バレッタ")
}

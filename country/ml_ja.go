//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Japanese, "マリ共和国")
	dataMali.RegisterOfficialName(xlanguage.Japanese, "マリ共和国")
	dataMali.RegisterCapital(xlanguage.Japanese, "バマコ")
}

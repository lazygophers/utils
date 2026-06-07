//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Japanese, "マダガスカル")
	dataMadagascar.RegisterOfficialName(xlanguage.Japanese, "マダガスカル共和国")
	dataMadagascar.RegisterCapital(xlanguage.Japanese, "アンタナナリボ")
}

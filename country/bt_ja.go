//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Japanese, "ブータン")
	dataBhutan.RegisterOfficialName(xlanguage.Japanese, "ブータン王国")
	dataBhutan.RegisterCapital(xlanguage.Japanese, "ティンプー")
}

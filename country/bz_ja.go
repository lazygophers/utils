//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Japanese, "ベリーズ")
	dataBelize.RegisterOfficialName(xlanguage.Japanese, "ベリーズ")
	dataBelize.RegisterCapital(xlanguage.Japanese, "ベルモパン")
}

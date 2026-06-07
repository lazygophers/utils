//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Japanese, "ブルキナファソ")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Japanese, "ブルキナファソ")
	dataBurkinaFaso.RegisterCapital(xlanguage.Japanese, "ワガドゥグー")
}

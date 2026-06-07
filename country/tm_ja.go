//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Japanese, "トルクメニスタン")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Japanese, "トルクメニスタン")
	dataTurkmenistan.RegisterCapital(xlanguage.Japanese, "アシガバート")
}

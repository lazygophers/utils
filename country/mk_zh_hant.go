//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.MustParse("zh-Hant"), "北馬其頓")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "北馬其頓共和國")
	dataNorthMacedonia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "史高比耶")
}

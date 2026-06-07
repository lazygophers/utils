//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.MustParse("zh-Hant"), "科威特")
	dataKuwait.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "科威特國")
	dataKuwait.RegisterCapital(xlanguage.MustParse("zh-Hant"), "科威特市")
}

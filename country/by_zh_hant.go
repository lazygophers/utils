//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.MustParse("zh-Hant"), "白俄羅斯")
	dataBelarus.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "白俄羅斯共和國")
	dataBelarus.RegisterCapital(xlanguage.MustParse("zh-Hant"), "明斯克")
}

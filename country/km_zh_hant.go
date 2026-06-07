//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.MustParse("zh-Hant"), "葛摩")
	dataComoros.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "葛摩聯盟")
	dataComoros.RegisterCapital(xlanguage.MustParse("zh-Hant"), "莫洛尼")
}

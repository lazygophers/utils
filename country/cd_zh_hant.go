//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.MustParse("zh-Hant"), "剛果民主共和國")
	dataDrCongo.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "剛果民主共和國")
	dataDrCongo.RegisterCapital(xlanguage.MustParse("zh-Hant"), "金夏沙")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.MustParse("zh-Hant"), "拉脫維亞")
	dataLatvia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "拉脫維亞共和國")
	dataLatvia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "里加")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.MustParse("zh-Hant"), "紐埃")
	dataNiue.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "紐埃")
	dataNiue.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿洛菲")
}

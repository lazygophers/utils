//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.MustParse("zh-Hant"), "塞爾維亞")
	dataSerbia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "塞爾維亞共和國")
	dataSerbia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "貝爾格勒")
}

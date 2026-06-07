//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.MustParse("zh-Hant"), "哥斯大黎加")
	dataCostaRica.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "哥斯大黎加共和國")
	dataCostaRica.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖荷西")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.MustParse("zh-Hant"), "千里達及托巴哥")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "千里達及托巴哥共和國")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.MustParse("zh-Hant"), "西班牙港")
}

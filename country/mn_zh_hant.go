//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.MustParse("zh-Hant"), "蒙古")
	dataMongolia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蒙古國")
	dataMongolia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "烏蘭巴托")
}

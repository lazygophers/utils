//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.MustParse("zh-Hant"), "獅子山")
	dataSierraLeone.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "獅子山共和國")
	dataSierraLeone.RegisterCapital(xlanguage.MustParse("zh-Hant"), "自由城")
}

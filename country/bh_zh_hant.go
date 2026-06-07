//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.MustParse("zh-Hant"), "巴林")
	dataBahrain.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴林王國")
	dataBahrain.RegisterCapital(xlanguage.MustParse("zh-Hant"), "麥納瑪")
}

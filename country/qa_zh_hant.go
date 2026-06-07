//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.MustParse("zh-Hant"), "卡達")
	dataQatar.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "卡達國")
	dataQatar.RegisterCapital(xlanguage.MustParse("zh-Hant"), "杜哈")
}

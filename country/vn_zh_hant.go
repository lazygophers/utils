//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.MustParse("zh-Hant"), "越南")
	dataVietnam.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "越南社會主義共和國")
	dataVietnam.RegisterCapital(xlanguage.MustParse("zh-Hant"), "河內")
}

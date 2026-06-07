//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.MustParse("zh-Hant"), "烏拉圭")
	dataUruguay.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "烏拉圭東方共和國")
	dataUruguay.RegisterCapital(xlanguage.MustParse("zh-Hant"), "蒙特維多")
}

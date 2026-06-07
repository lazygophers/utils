//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.MustParse("zh-Hant"), "聖文森及格瑞那丁")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖文森及格瑞那丁")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.MustParse("zh-Hant"), "金石城")
}

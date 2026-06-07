//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.MustParse("zh-Hant"), "蒙哲臘")
	dataMontserrat.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蒙哲臘")
	dataMontserrat.RegisterCapital(xlanguage.MustParse("zh-Hant"), "普利茅斯")
}

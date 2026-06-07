//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.MustParse("zh-Hant"), "法屬聖馬丁")
	dataSaintMartin.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖馬丁集體")
	dataSaintMartin.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬里戈特")
}

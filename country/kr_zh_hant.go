//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.MustParse("zh-Hant"), "南韓")
	dataSouthKorea.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "大韓民國")
	dataSouthKorea.RegisterCapital(xlanguage.MustParse("zh-Hant"), "首爾")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.MustParse("zh-Hant"), "馬利")
	dataMali.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬利共和國")
	dataMali.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴馬科")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.MustParse("zh-Hant"), "莫三比克")
	dataMozambique.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "莫三比克共和國")
	dataMozambique.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬普托")
}

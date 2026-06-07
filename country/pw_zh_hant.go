//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.MustParse("zh-Hant"), "帛琉")
	dataPalau.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "帛琉共和國")
	dataPalau.RegisterCapital(xlanguage.MustParse("zh-Hant"), "恩吉魯穆德")
}

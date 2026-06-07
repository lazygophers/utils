//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.MustParse("zh-Hant"), "蒙特內哥羅")
	dataMontenegro.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蒙特內哥羅")
	dataMontenegro.RegisterCapital(xlanguage.MustParse("zh-Hant"), "波多里查")
}

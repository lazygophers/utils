//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.MustParse("zh-Hant"), "巴勒斯坦")
	dataPalestine.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "巴勒斯坦國")
	dataPalestine.RegisterCapital(xlanguage.MustParse("zh-Hant"), "拉姆安拉")
}

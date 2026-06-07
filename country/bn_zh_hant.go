//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.MustParse("zh-Hant"), "汶萊")
	dataBrunei.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "汶萊和平之國")
	dataBrunei.RegisterCapital(xlanguage.MustParse("zh-Hant"), "斯里巴卡旺市")
}

//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.MustParse("zh-Hant"), "聖赫倫那、亞森欣與崔斯坦達庫尼亞")
	dataSaintHelena.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖赫倫那、亞森欣與崔斯坦達庫尼亞")
	dataSaintHelena.RegisterCapital(xlanguage.MustParse("zh-Hant"), "詹姆斯敦")
}

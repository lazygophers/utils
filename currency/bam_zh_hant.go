//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bam.RegisterName(xlanguage.MustParse("zh-Hant"), "波士尼亞與赫塞哥維納可兌換馬克")
}

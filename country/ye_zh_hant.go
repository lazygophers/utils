//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.MustParse("zh-Hant"), "葉門")
	dataYemen.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "葉門共和國")
	dataYemen.RegisterCapital(xlanguage.MustParse("zh-Hant"), "沙那")
}

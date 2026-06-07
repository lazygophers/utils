//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_bm || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.MustParse("zh-Hant"), "百慕達")
	dataBermuda.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "百慕達")
	dataBermuda.RegisterCapital(xlanguage.MustParse("zh-Hant"), "漢密爾頓")
}

//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_cg || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.MustParse("zh-Hant"), "剛果共和國")
	dataCongo.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "剛果共和國")
	dataCongo.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布拉薩市")
}

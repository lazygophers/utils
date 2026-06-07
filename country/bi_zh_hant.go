//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.MustParse("zh-Hant"), "蒲隆地")
	dataBurundi.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蒲隆地共和國")
	dataBurundi.RegisterCapital(xlanguage.MustParse("zh-Hant"), "基特加")
}

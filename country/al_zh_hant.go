//go:build (lang_zh_hant || lang_all) && (country_al || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.MustParse("zh-Hant"), "阿爾巴尼亞")
	dataAlbania.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿爾巴尼亞共和國")
	dataAlbania.RegisterCapital(xlanguage.MustParse("zh-Hant"), "地拉那")
}

//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_my || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.MustParse("zh-Hant"), "馬來西亞")
	dataMalaysia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬來西亞")
	dataMalaysia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "吉隆坡")
}

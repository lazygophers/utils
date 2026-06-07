//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_bf || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.MustParse("zh-Hant"), "布吉納法索")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "布吉納法索")
	dataBurkinaFaso.RegisterCapital(xlanguage.MustParse("zh-Hant"), "瓦加杜古")
}

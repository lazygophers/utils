//go:build (lang_zh_hant || lang_all) && (country_all || country_micronesia || country_mp || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "北馬利安納群島")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "北馬利安納群島自由邦")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "塞班島")
}

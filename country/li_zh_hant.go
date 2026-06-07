//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_li || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.MustParse("zh-Hant"), "列支敦斯登")
	dataLiechtenstein.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "列支敦斯登親王國")
	dataLiechtenstein.RegisterCapital(xlanguage.MustParse("zh-Hant"), "瓦都茲")
}

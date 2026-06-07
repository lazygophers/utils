//go:build (lang_zh_hant || lang_all) && (country_all || country_antarctic || country_gs)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.MustParse("zh-Hant"), "南喬治亞和南桑威奇群島")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "南喬治亞和南桑威奇群島")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.MustParse("zh-Hant"), "愛德華國王角")
}

//go:build (lang_zh_hant || lang_all) && (country_all || country_oceania || country_polynesia || country_wf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.MustParse("zh-Hant"), "瓦利斯和富圖納")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "瓦利斯和富圖納群島領地")
	dataWallisAndFutuna.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬塔烏圖")
}

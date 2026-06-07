//go:build country_all || country_oceania || country_polynesia || country_wf

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Chinese, "瓦利斯和富图纳")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Chinese, "瓦利斯和富图纳群岛")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Chinese, "马塔乌图")
}

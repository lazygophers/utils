//go:build country_af || country_all || country_asia || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Chinese, "阿富汗")
	dataAfghanistan.RegisterOfficialName(xlanguage.Chinese, "阿富汗伊斯兰酋长国")
	dataAfghanistan.RegisterCapital(xlanguage.Chinese, "喀布尔")
}

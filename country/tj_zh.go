//go:build country_all || country_asia || country_central_asia || country_tj

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Chinese, "塔吉克斯坦")
	dataTajikistan.RegisterOfficialName(xlanguage.Chinese, "塔吉克斯坦共和国")
	dataTajikistan.RegisterCapital(xlanguage.Chinese, "杜尚别")
}

//go:build country_africa || country_all || country_eastern_africa || country_mg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Chinese, "马达加斯加")
	dataMadagascar.RegisterOfficialName(xlanguage.Chinese, "马达加斯加共和国")
	dataMadagascar.RegisterCapital(xlanguage.Chinese, "塔那那利佛")
}

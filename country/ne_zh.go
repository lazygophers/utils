//go:build country_africa || country_all || country_ne || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.Chinese, "尼日尔")
	dataNiger.RegisterOfficialName(xlanguage.Chinese, "尼日尔共和国")
	dataNiger.RegisterCapital(xlanguage.Chinese, "尼亚美")
}

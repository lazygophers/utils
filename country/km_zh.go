//go:build country_africa || country_all || country_eastern_africa || country_km

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Chinese, "科摩罗")
	dataComoros.RegisterOfficialName(xlanguage.Chinese, "科摩罗联盟")
	dataComoros.RegisterCapital(xlanguage.Chinese, "莫罗尼")
}

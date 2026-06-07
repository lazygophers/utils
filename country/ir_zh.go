//go:build country_all || country_asia || country_ir || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Chinese, "伊朗")
	dataIran.RegisterOfficialName(xlanguage.Chinese, "伊朗伊斯兰共和国")
	dataIran.RegisterCapital(xlanguage.Chinese, "德黑兰")
}

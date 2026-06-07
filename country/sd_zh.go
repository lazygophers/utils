//go:build country_africa || country_all || country_northern_africa || country_sd

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Chinese, "苏丹")
	dataSudan.RegisterOfficialName(xlanguage.Chinese, "苏丹共和国")
	dataSudan.RegisterCapital(xlanguage.Chinese, "喀土穆")
}

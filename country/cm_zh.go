//go:build country_africa || country_all || country_cm || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Chinese, "喀麦隆")
	dataCameroon.RegisterOfficialName(xlanguage.Chinese, "喀麦隆共和国")
	dataCameroon.RegisterCapital(xlanguage.Chinese, "雅温得")
}

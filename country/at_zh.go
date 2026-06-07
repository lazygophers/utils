//go:build country_all || country_at || country_europe || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Chinese, "奥地利")
	dataAustria.RegisterOfficialName(xlanguage.Chinese, "奥地利共和国")
	dataAustria.RegisterCapital(xlanguage.Chinese, "维也纳")
}

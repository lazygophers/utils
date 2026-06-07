//go:build country_all || country_es || country_europe || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.Chinese, "西班牙")
	dataSpain.RegisterOfficialName(xlanguage.Chinese, "西班牙王国")
	dataSpain.RegisterCapital(xlanguage.Chinese, "马德里")
}

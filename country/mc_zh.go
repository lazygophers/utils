//go:build country_all || country_europe || country_mc || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.Chinese, "摩纳哥")
	dataMonaco.RegisterOfficialName(xlanguage.Chinese, "摩纳哥公国")
	dataMonaco.RegisterCapital(xlanguage.Chinese, "摩纳哥")
}

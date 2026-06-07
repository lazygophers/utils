//go:build country_africa || country_all || country_middle_africa || country_st

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Chinese, "圣多美和普林西比")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Chinese, "圣多美和普林西比民主共和国")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Chinese, "圣多美")
}

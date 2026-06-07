//go:build country_africa || country_all || country_eastern_africa || country_ug

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Chinese, "乌干达")
	dataUganda.RegisterOfficialName(xlanguage.Chinese, "乌干达共和国")
	dataUganda.RegisterCapital(xlanguage.Chinese, "坎帕拉")
}

//go:build country_all || country_by || country_eastern_europe || country_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Chinese, "白俄罗斯")
	dataBelarus.RegisterOfficialName(xlanguage.Chinese, "白俄罗斯共和国")
	dataBelarus.RegisterCapital(xlanguage.Chinese, "明斯克")
}

//go:build country_africa || country_all || country_eastern_africa || country_mw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Chinese, "马拉维")
	dataMalawi.RegisterOfficialName(xlanguage.Chinese, "马拉维共和国")
	dataMalawi.RegisterCapital(xlanguage.Chinese, "利隆圭")
}

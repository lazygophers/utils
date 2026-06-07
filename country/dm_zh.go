//go:build country_all || country_americas || country_caribbean || country_dm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Chinese, "多米尼克")
	dataDominica.RegisterOfficialName(xlanguage.Chinese, "多米尼克国")
	dataDominica.RegisterCapital(xlanguage.Chinese, "罗索")
}

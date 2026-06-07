//go:build country_africa || country_all || country_eastern_africa || country_tz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Chinese, "坦桑尼亚")
	dataTanzania.RegisterOfficialName(xlanguage.Chinese, "坦桑尼亚联合共和国")
	dataTanzania.RegisterCapital(xlanguage.Chinese, "多多马")
}

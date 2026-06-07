//go:build country_africa || country_all || country_cf || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.Chinese, "中非共和国")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.Chinese, "中非共和国")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.Chinese, "班吉")
}

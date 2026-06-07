//go:build country_all || country_europe || country_southern_europe || country_va

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Chinese, "梵蒂冈")
	dataVaticanCity.RegisterOfficialName(xlanguage.Chinese, "梵蒂冈城国")
	dataVaticanCity.RegisterCapital(xlanguage.Chinese, "梵蒂冈城")
}

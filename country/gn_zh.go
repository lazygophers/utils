//go:build country_africa || country_all || country_gn || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Chinese, "几内亚")
	dataGuinea.RegisterOfficialName(xlanguage.Chinese, "几内亚共和国")
	dataGuinea.RegisterCapital(xlanguage.Chinese, "科纳克里")
}

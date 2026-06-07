//go:build country_africa || country_all || country_ao || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Chinese, "安哥拉")
	dataAngola.RegisterOfficialName(xlanguage.Chinese, "安哥拉共和国")
	dataAngola.RegisterCapital(xlanguage.Chinese, "罗安达")
}

//go:build country_all || country_americas || country_caribbean || country_ht

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.Chinese, "海地")
	dataHaiti.RegisterOfficialName(xlanguage.Chinese, "海地共和国")
	dataHaiti.RegisterCapital(xlanguage.Chinese, "太子港")
}

//go:build country_all || country_asia || country_iq || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.Chinese, "伊拉克")
	dataIraq.RegisterOfficialName(xlanguage.Chinese, "伊拉克共和国")
	dataIraq.RegisterCapital(xlanguage.Chinese, "巴格达")
}

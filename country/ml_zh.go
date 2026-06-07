//go:build country_africa || country_all || country_ml || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Chinese, "马里")
	dataMali.RegisterOfficialName(xlanguage.Chinese, "马里共和国")
	dataMali.RegisterCapital(xlanguage.Chinese, "巴马科")
}

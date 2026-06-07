//go:build country_africa || country_all || country_cv || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.Chinese, "佛得角")
	dataCaboVerde.RegisterOfficialName(xlanguage.Chinese, "佛得角共和国")
	dataCaboVerde.RegisterCapital(xlanguage.Chinese, "普拉亚")
}

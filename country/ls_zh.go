//go:build country_africa || country_all || country_ls || country_southern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Chinese, "莱索托")
	dataLesotho.RegisterOfficialName(xlanguage.Chinese, "莱索托王国")
	dataLesotho.RegisterCapital(xlanguage.Chinese, "马塞卢")
}

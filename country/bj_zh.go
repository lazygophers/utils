//go:build country_africa || country_all || country_bj || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.Chinese, "贝宁")
	dataBenin.RegisterOfficialName(xlanguage.Chinese, "贝宁共和国")
	dataBenin.RegisterCapital(xlanguage.Chinese, "波多诺伏")
}

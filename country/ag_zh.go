//go:build country_ag || country_all || country_americas || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Chinese, "安提瓜和巴布达")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Chinese, "安提瓜和巴布达")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Chinese, "圣约翰")
}

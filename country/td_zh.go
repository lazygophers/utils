//go:build country_africa || country_all || country_middle_africa || country_td

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Chinese, "乍得")
	dataChad.RegisterOfficialName(xlanguage.Chinese, "乍得共和国")
	dataChad.RegisterCapital(xlanguage.Chinese, "恩贾梅纳")
}

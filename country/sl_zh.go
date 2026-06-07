//go:build country_africa || country_all || country_sl || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Chinese, "塞拉利昂")
	dataSierraLeone.RegisterOfficialName(xlanguage.Chinese, "塞拉利昂共和国")
	dataSierraLeone.RegisterCapital(xlanguage.Chinese, "弗里敦")
}

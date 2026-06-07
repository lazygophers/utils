//go:build country_africa || country_all || country_tg || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Chinese, "多哥")
	dataTogo.RegisterOfficialName(xlanguage.Chinese, "多哥共和国")
	dataTogo.RegisterCapital(xlanguage.Chinese, "洛美")
}

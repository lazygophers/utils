//go:build country_all || country_asia || country_kh || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Chinese, "柬埔寨")
	dataCambodia.RegisterOfficialName(xlanguage.Chinese, "柬埔寨王国")
	dataCambodia.RegisterCapital(xlanguage.Chinese, "金边")
}

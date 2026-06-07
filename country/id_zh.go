//go:build country_all || country_asia || country_id || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.Chinese, "印度尼西亚")
	dataIndonesia.RegisterOfficialName(xlanguage.Chinese, "印度尼西亚共和国")
	dataIndonesia.RegisterCapital(xlanguage.Chinese, "雅加达")
}

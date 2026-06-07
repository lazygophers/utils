//go:build country_africa || country_all || country_eastern_africa || country_rw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Chinese, "卢旺达")
	dataRwanda.RegisterOfficialName(xlanguage.Chinese, "卢旺达共和国")
	dataRwanda.RegisterCapital(xlanguage.Chinese, "基加利")
}

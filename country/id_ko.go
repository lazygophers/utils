//go:build (lang_ko || lang_all) && (country_all || country_asia || country_id || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.Korean, "인도네시아")
	dataIndonesia.RegisterOfficialName(xlanguage.Korean, "인도네시아 공화국")
	dataIndonesia.RegisterCapital(xlanguage.Korean, "자카르타")
}

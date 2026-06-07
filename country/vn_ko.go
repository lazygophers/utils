//go:build (lang_ko || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_vn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.Korean, "베트남")
	dataVietnam.RegisterOfficialName(xlanguage.Korean, "베트남 사회주의 공화국")
	dataVietnam.RegisterCapital(xlanguage.Korean, "하노이")
}

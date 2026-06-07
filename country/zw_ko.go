//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Korean, "짐바브웨")
	dataZimbabwe.RegisterOfficialName(xlanguage.Korean, "짐바브웨 공화국")
	dataZimbabwe.RegisterCapital(xlanguage.Korean, "하라레")
}

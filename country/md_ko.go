//go:build (lang_ko || lang_all) && (country_all || country_eastern_europe || country_europe || country_md)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.Korean, "몰도바")
	dataMoldova.RegisterOfficialName(xlanguage.Korean, "몰도바 공화국")
	dataMoldova.RegisterCapital(xlanguage.Korean, "키시너우")
}

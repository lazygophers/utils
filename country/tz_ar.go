//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Arabic, "تنزانيا")
	dataTanzania.RegisterOfficialName(xlanguage.Arabic, "جمهورية تنزانيا الاتحادية")
	dataTanzania.RegisterCapital(xlanguage.Arabic, "دودوما")
}

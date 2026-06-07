//go:build (lang_ar || lang_all) && (country_all || country_europe || country_lu || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Arabic, "لوكسمبورغ")
	dataLuxembourg.RegisterOfficialName(xlanguage.Arabic, "دوقية لوكسمبورغ الكبرى")
	dataLuxembourg.RegisterCapital(xlanguage.Arabic, "مدينة لوكسمبورغ")
}

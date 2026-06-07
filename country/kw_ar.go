//go:build country_all || country_asia || country_kw || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Arabic, "الكويت")
	dataKuwait.RegisterOfficialName(xlanguage.Arabic, "دولة الكويت")
	dataKuwait.RegisterCapital(xlanguage.Arabic, "مدينة الكويت")
}

//go:build (lang_ar || lang_all) && (country_all || country_europe || country_me || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.Arabic, "الجبل الأسود")
	dataMontenegro.RegisterOfficialName(xlanguage.Arabic, "الجبل الأسود")
	dataMontenegro.RegisterCapital(xlanguage.Arabic, "بودغوريتسا")
}

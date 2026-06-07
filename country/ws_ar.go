//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_polynesia || country_ws)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.Arabic, "ساموا")
	dataSamoa.RegisterOfficialName(xlanguage.Arabic, "دولة ساموا المستقلة")
	dataSamoa.RegisterCapital(xlanguage.Arabic, "آبيا")
}

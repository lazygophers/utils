//go:build (lang_ar || lang_all) && (country_all || country_europe || country_it || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.Arabic, "إيطاليا")
	dataItaly.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الإيطالية")
	dataItaly.RegisterCapital(xlanguage.Arabic, "روما")
}

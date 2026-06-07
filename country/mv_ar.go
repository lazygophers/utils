//go:build (lang_ar || lang_all) && (country_all || country_asia || country_mv || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Arabic, "المالديف")
	dataMaldives.RegisterOfficialName(xlanguage.Arabic, "جمهورية المالديف")
	dataMaldives.RegisterCapital(xlanguage.Arabic, "ماليه")
}

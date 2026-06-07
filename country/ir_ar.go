//go:build (lang_ar || lang_all) && (country_all || country_asia || country_ir || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Arabic, "إيران")
	dataIran.RegisterOfficialName(xlanguage.Arabic, "جمهورية إيران الإسلامية")
	dataIran.RegisterCapital(xlanguage.Arabic, "طهران")
}

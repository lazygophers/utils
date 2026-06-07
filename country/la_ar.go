//go:build (lang_ar || lang_all) && (country_all || country_asia || country_la || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Arabic, "لاوس")
	dataLaos.RegisterOfficialName(xlanguage.Arabic, "جمهورية لاوس الديمقراطية الشعبية")
	dataLaos.RegisterCapital(xlanguage.Arabic, "فيينتيان")
}

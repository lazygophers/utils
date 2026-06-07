//go:build (lang_ar || lang_all) && (country_af || country_all || country_asia || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Arabic, "أفغانستان")
	dataAfghanistan.RegisterOfficialName(xlanguage.Arabic, "إمارة أفغانستان الإسلامية")
	dataAfghanistan.RegisterCapital(xlanguage.Arabic, "كابول")
}

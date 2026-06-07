//go:build (lang_ar || lang_all) && (country_all || country_asia || country_pk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Arabic, "باكستان")
	dataPakistan.RegisterOfficialName(xlanguage.Arabic, "جمهورية باكستان الإسلامية")
	dataPakistan.RegisterCapital(xlanguage.Arabic, "إسلام أباد")
}

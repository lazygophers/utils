//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_pf || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Arabic, "بولينزيا الفرنسية")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Arabic, "بولينزيا الفرنسية")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Arabic, "بابيتي")
}

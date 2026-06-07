//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Arabic, "بولينزيا الفرنسية")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Arabic, "بولينزيا الفرنسية")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Arabic, "بابيتي")
}

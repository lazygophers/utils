//go:build (lang_ja || lang_all) && (country_all || country_oceania || country_pf || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Japanese, "フランス領ポリネシア")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Japanese, "フランス領ポリネシア")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Japanese, "パペーテ")
}

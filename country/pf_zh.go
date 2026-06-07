//go:build country_all || country_oceania || country_pf || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Chinese, "法属波利尼西亚")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Chinese, "法属波利尼西亚")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Chinese, "帕皮提")
}

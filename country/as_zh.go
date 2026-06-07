//go:build country_all || country_as || country_oceania || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.Chinese, "美属萨摩亚")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.Chinese, "美属萨摩亚领地")
	dataAmericanSamoa.RegisterCapital(xlanguage.Chinese, "帕果帕果")
}

//go:build country_all || country_au || country_australia_and_new_zealand || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Chinese, "澳大利亚")
	dataAustralia.RegisterOfficialName(xlanguage.Chinese, "澳大利亚联邦")
	dataAustralia.RegisterCapital(xlanguage.Chinese, "堪培拉")
}

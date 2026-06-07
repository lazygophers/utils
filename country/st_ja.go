//go:build (lang_ja || lang_all) && (country_africa || country_all || country_middle_africa || country_st)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Japanese, "サントメ・プリンシペ")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Japanese, "サントメ・プリンシペ民主共和国")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Japanese, "サントメ")
}

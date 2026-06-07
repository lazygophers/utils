//go:build (lang_ar || lang_all) && (country_africa || country_all || country_middle_africa || country_st)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Arabic, "ساو تومي وبرينسيب")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Arabic, "جمهورية ساو تومي وبرينسيب الديمقراطية")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Arabic, "ساو تومي")
}

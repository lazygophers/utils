//go:build country_all || country_asia || country_sa || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.Arabic, "المملكة العربية السعودية")
	dataSaudiArabia.RegisterOfficialName(xlanguage.Arabic, "المملكة العربية السعودية")
	dataSaudiArabia.RegisterCapital(xlanguage.Arabic, "الرياض")
}

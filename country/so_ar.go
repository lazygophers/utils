//go:build country_africa || country_all || country_eastern_africa || country_so

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Arabic, "الصومال")
	dataSomalia.RegisterOfficialName(xlanguage.Arabic, "جمهورية الصومال الاتحادية")
	dataSomalia.RegisterCapital(xlanguage.Arabic, "مقديشو")
}

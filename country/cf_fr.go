//go:build country_africa || country_all || country_cf || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.French, "République centrafricaine")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.French, "République centrafricaine")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.French, "Bangui")
}

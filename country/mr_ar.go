//go:build country_africa || country_all || country_mr || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.Arabic, "موريتانيا")
	dataMauritania.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الإسلامية الموريتانية")
	dataMauritania.RegisterCapital(xlanguage.Arabic, "نواكشوط")
}

//go:build (lang_es || lang_all) && (country_africa || country_all || country_mr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.Spanish, "Mauritania")
	dataMauritania.RegisterOfficialName(xlanguage.Spanish, "República Islámica de Mauritania")
	dataMauritania.RegisterCapital(xlanguage.Spanish, "Nuakchot")
}

//go:build (lang_fr || lang_all) && (country_africa || country_all || country_mr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.French, "Mauritanie")
	dataMauritania.RegisterOfficialName(xlanguage.French, "République islamique de Mauritanie")
	dataMauritania.RegisterCapital(xlanguage.French, "Nouakchott")
}

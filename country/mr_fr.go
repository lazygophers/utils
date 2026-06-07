//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.French, "Mauritanie")
	dataMauritania.RegisterOfficialName(xlanguage.French, "République islamique de Mauritanie")
	dataMauritania.RegisterCapital(xlanguage.French, "Nouakchott")
}

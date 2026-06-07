//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.Russian, "Багамские Острова")
	dataBahamas.RegisterOfficialName(xlanguage.Russian, "Содружество Багамских Островов")
	dataBahamas.RegisterCapital(xlanguage.Russian, "Нассау")
}

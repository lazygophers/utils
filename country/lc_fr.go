//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.French, "Sainte-Lucie")
	dataSaintLucia.RegisterOfficialName(xlanguage.French, "Sainte-Lucie")
	dataSaintLucia.RegisterCapital(xlanguage.French, "Castries")
}

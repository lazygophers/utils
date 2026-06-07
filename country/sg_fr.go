//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.French, "Singapour")
	dataSingapore.RegisterOfficialName(xlanguage.French, "République de Singapour")
	dataSingapore.RegisterCapital(xlanguage.French, "Singapour")
}

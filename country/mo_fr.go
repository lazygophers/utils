//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.French, "Macao")
	dataMacao.RegisterOfficialName(xlanguage.French, "Région administrative spéciale de Macao")
	dataMacao.RegisterCapital(xlanguage.French, "Macao")
}

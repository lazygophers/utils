//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Spanish, "Yibuti")
	dataDjibouti.RegisterOfficialName(xlanguage.Spanish, "República de Yibuti")
	dataDjibouti.RegisterCapital(xlanguage.Spanish, "Yibuti")
}

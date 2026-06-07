//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Spanish, "Dinamarca")
	dataDenmark.RegisterOfficialName(xlanguage.Spanish, "Reino de Dinamarca")
	dataDenmark.RegisterCapital(xlanguage.Spanish, "Copenhague")
}

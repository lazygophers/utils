//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.Spanish, "Grecia")
	dataGreece.RegisterOfficialName(xlanguage.Spanish, "República Helénica")
	dataGreece.RegisterCapital(xlanguage.Spanish, "Atenas")
}

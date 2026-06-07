//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Spanish, "Guyana")
	dataGuyana.RegisterOfficialName(xlanguage.Spanish, "República Cooperativa de Guyana")
	dataGuyana.RegisterCapital(xlanguage.Spanish, "Georgetown")
}

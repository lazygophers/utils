//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.Spanish, "China")
	dataChina.RegisterOfficialName(xlanguage.Spanish, "República Popular China")
	dataChina.RegisterCapital(xlanguage.Spanish, "Pekín")
}

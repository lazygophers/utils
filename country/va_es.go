//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Spanish, "Ciudad del Vaticano")
	dataVaticanCity.RegisterOfficialName(xlanguage.Spanish, "Estado de la Ciudad del Vaticano")
	dataVaticanCity.RegisterCapital(xlanguage.Spanish, "Ciudad del Vaticano")
}

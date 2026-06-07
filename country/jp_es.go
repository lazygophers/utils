//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.Spanish, "Japón")
	dataJapan.RegisterOfficialName(xlanguage.Spanish, "Estado del Japón")
	dataJapan.RegisterCapital(xlanguage.Spanish, "Tokio")
}

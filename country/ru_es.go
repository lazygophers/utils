//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.Spanish, "Rusia")
	dataRussia.RegisterOfficialName(xlanguage.Spanish, "Federación de Rusia")
	dataRussia.RegisterCapital(xlanguage.Spanish, "Moscú")
}

//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Spanish, "Tonga")
	dataTonga.RegisterOfficialName(xlanguage.Spanish, "Reino de Tonga")
	dataTonga.RegisterCapital(xlanguage.Spanish, "Nukualofa")
}

//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Spanish, "Islas Ultramarinas Menores de Estados Unidos")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Ultramarinas Menores de Estados Unidos")
}

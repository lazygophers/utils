//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Spanish, "Brunéi")
	dataBrunei.RegisterOfficialName(xlanguage.Spanish, "Nación de Brunéi, Morada de la Paz")
	dataBrunei.RegisterCapital(xlanguage.Spanish, "Bandar Seri Begawan")
}

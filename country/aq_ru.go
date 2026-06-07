//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Russian, "Антарктида")
	dataAntarctica.RegisterOfficialName(xlanguage.Russian, "Антарктида")
}

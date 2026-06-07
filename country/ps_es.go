//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Spanish, "Palestina")
	dataPalestine.RegisterOfficialName(xlanguage.Spanish, "Estado de Palestina")
	dataPalestine.RegisterCapital(xlanguage.Spanish, "Jerusalén Este")
}

//go:build (lang_es || lang_all) && (country_all || country_asia || country_ps || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Spanish, "Palestina")
	dataPalestine.RegisterOfficialName(xlanguage.Spanish, "Estado de Palestina")
	dataPalestine.RegisterCapital(xlanguage.Spanish, "Jerusalén Este")
}

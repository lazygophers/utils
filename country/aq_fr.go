//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.French, "Antarctique")
	dataAntarctica.RegisterOfficialName(xlanguage.French, "Antarctique")
}

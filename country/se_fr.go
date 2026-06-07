//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.French, "Suède")
	dataSweden.RegisterOfficialName(xlanguage.French, "Royaume de Suède")
	dataSweden.RegisterCapital(xlanguage.French, "Stockholm")
}

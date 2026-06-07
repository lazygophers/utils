//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Spanish, "Túnez")
	dataTunisia.RegisterOfficialName(xlanguage.Spanish, "República Tunecina")
	dataTunisia.RegisterCapital(xlanguage.Spanish, "Túnez")
}

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Eur.RegisterName(xlanguage.English, "Euro")
}

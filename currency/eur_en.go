package currency

import xlanguage "golang.org/x/text/language"

func init() {
	EUR.RegisterName(xlanguage.English, "Euro")
}

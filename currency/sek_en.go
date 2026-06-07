package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sek.RegisterName(xlanguage.English, "Swedish Krona")
}

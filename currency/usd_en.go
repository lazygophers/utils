package currency

import xlanguage "golang.org/x/text/language"

func init() {
	USD.RegisterName(xlanguage.English, "US Dollar")
}

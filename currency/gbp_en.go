package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GBP.RegisterName(xlanguage.English, "Pound Sterling")
}

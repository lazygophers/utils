package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.English, "Wallis and Futuna")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.English, "Territory of the Wallis and Futuna Islands")
	dataWallisAndFutuna.RegisterCapital(xlanguage.English, "Mata-Utu")
}

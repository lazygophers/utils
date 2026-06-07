package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.French, "Wallis-et-Futuna")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.French, "Territoire des îles Wallis-et-Futuna")
	dataWallisAndFutuna.RegisterCapital(xlanguage.French, "Mata-Utu")
}

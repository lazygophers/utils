//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.French, "Géorgie")
	dataGeorgia.RegisterOfficialName(xlanguage.French, "Géorgie")
	dataGeorgia.RegisterCapital(xlanguage.French, "Tbilissi")
}

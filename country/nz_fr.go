//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.French, "Nouvelle-Zélande")
	dataNewZealand.RegisterOfficialName(xlanguage.French, "Nouvelle-Zélande")
	dataNewZealand.RegisterCapital(xlanguage.French, "Wellington")
}

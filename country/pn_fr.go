//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.French, "Îles Pitcairn")
	dataPitcairn.RegisterOfficialName(xlanguage.French, "Îles Pitcairn, Henderson, Ducie et Oeno")
	dataPitcairn.RegisterCapital(xlanguage.French, "Adamstown")
}

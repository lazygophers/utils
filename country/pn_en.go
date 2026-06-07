//go:build country_all || country_oceania || country_pn || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.English, "Pitcairn Islands")
	dataPitcairn.RegisterOfficialName(xlanguage.English, "Pitcairn, Henderson, Ducie and Oeno Islands")
	dataPitcairn.RegisterCapital(xlanguage.English, "Adamstown")
}

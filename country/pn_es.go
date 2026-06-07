//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Spanish, "Islas Pitcairn")
	dataPitcairn.RegisterOfficialName(xlanguage.Spanish, "Islas Pitcairn")
	dataPitcairn.RegisterCapital(xlanguage.Spanish, "Adamstown")
}

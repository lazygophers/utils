//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.French, "Kiribati")
	dataKiribati.RegisterOfficialName(xlanguage.French, "République de Kiribati")
	dataKiribati.RegisterCapital(xlanguage.French, "Tarawa-Sud")
}

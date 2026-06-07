//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.French, "Corée du Sud")
	dataSouthKorea.RegisterOfficialName(xlanguage.French, "République de Corée")
	dataSouthKorea.RegisterCapital(xlanguage.French, "Séoul")
}

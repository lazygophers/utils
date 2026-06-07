//go:build country_africa || country_all || country_sn || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.French, "Sénégal")
	dataSenegal.RegisterOfficialName(xlanguage.French, "République du Sénégal")
	dataSenegal.RegisterCapital(xlanguage.French, "Dakar")
}

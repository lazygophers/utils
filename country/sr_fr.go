//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.French, "Suriname")
	dataSuriname.RegisterOfficialName(xlanguage.French, "République du Suriname")
	dataSuriname.RegisterCapital(xlanguage.French, "Paramaribo")
}

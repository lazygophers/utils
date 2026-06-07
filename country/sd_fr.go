//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.French, "Soudan")
	dataSudan.RegisterOfficialName(xlanguage.French, "République du Soudan")
	dataSudan.RegisterCapital(xlanguage.French, "Khartoum")
}

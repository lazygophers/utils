//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.French, "Taïwan")
	dataTaiwan.RegisterOfficialName(xlanguage.French, "République de Chine")
	dataTaiwan.RegisterCapital(xlanguage.French, "Taipei")
}

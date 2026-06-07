//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.French, "Biélorussie")
	dataBelarus.RegisterOfficialName(xlanguage.French, "République de Biélorussie")
	dataBelarus.RegisterCapital(xlanguage.French, "Minsk")
}

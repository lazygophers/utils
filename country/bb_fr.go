//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.French, "Barbade")
	dataBarbados.RegisterOfficialName(xlanguage.French, "Barbade")
	dataBarbados.RegisterCapital(xlanguage.French, "Bridgetown")
}

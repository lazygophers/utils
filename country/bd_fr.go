//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.French, "Bangladesh")
	dataBangladesh.RegisterOfficialName(xlanguage.French, "République populaire du Bangladesh")
	dataBangladesh.RegisterCapital(xlanguage.French, "Dacca")
}

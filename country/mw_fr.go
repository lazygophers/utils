//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.French, "Malawi")
	dataMalawi.RegisterOfficialName(xlanguage.French, "République du Malawi")
	dataMalawi.RegisterCapital(xlanguage.French, "Lilongwe")
}

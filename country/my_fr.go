//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.French, "Malaisie")
	dataMalaysia.RegisterOfficialName(xlanguage.French, "Malaisie")
	dataMalaysia.RegisterCapital(xlanguage.French, "Kuala Lumpur")
}

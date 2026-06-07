//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.French, "Guatemala")
	dataGuatemala.RegisterOfficialName(xlanguage.French, "République du Guatemala")
	dataGuatemala.RegisterCapital(xlanguage.French, "Guatemala")
}

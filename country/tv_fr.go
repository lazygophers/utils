//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.French, "Tuvalu")
	dataTuvalu.RegisterOfficialName(xlanguage.French, "Tuvalu")
	dataTuvalu.RegisterCapital(xlanguage.French, "Funafuti")
}

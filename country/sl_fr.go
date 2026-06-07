//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.French, "Sierra Leone")
	dataSierraLeone.RegisterOfficialName(xlanguage.French, "République de Sierra Leone")
	dataSierraLeone.RegisterCapital(xlanguage.French, "Freetown")
}

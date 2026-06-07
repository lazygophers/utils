//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.French, "Aruba")
	dataAruba.RegisterOfficialName(xlanguage.French, "Aruba")
	dataAruba.RegisterCapital(xlanguage.French, "Oranjestad")
}

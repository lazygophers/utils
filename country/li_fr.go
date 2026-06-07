//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.French, "Liechtenstein")
	dataLiechtenstein.RegisterOfficialName(xlanguage.French, "Principauté de Liechtenstein")
	dataLiechtenstein.RegisterCapital(xlanguage.French, "Vaduz")
}

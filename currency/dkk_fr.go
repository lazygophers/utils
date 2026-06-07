//go:build lang_fr || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dkk.RegisterName(xlanguage.French, "Couronne danoise")
}

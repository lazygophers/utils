//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gmd.RegisterName(xlanguage.Spanish, "Dalasi")
}

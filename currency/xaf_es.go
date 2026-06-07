//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xaf.RegisterName(xlanguage.Spanish, "Franco CFA de África Central")
}

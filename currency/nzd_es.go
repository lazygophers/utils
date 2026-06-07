//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nzd.RegisterName(xlanguage.Spanish, "Dólar neozelandés")
}

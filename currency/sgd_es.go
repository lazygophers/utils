//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sgd.RegisterName(xlanguage.Spanish, "Dólar de Singapur")
}

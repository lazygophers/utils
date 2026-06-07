//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Twd.RegisterName(xlanguage.Spanish, "Nuevo dólar taiwanés")
}

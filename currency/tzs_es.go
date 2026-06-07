//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tzs.RegisterName(xlanguage.Spanish, "Chelín tanzano")
}

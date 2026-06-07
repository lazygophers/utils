//go:build lang_es || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tnd.RegisterName(xlanguage.Spanish, "Dinar tunecino")
}

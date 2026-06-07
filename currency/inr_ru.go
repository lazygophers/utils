//go:build lang_ru || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Inr.RegisterName(xlanguage.Russian, "Индийская рупия")
}

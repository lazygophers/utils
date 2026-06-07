//go:build lang_fr || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pen.RegisterName(xlanguage.French, "Sol péruvien")
}

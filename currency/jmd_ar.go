//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Jmd.RegisterName(xlanguage.Arabic, "دولار جامايكي")
}

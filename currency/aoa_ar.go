//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Aoa.RegisterName(xlanguage.Arabic, "كوانزا")
}

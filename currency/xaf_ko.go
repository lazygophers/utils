//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xaf.RegisterName(xlanguage.Korean, "중앙아프리카 CFA 프랑")
}

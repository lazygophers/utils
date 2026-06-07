//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Korean, "덴마크")
	dataDenmark.RegisterOfficialName(xlanguage.Korean, "덴마크 왕국")
	dataDenmark.RegisterCapital(xlanguage.Korean, "코펜하겐")
}

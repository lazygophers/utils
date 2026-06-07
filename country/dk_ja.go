//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Japanese, "デンマーク")
	dataDenmark.RegisterOfficialName(xlanguage.Japanese, "デンマーク王国")
	dataDenmark.RegisterCapital(xlanguage.Japanese, "コペンハーゲン")
}

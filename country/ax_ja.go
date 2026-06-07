//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Japanese, "オーランド諸島")
	dataAlandIslands.RegisterOfficialName(xlanguage.Japanese, "オーランド諸島")
	dataAlandIslands.RegisterCapital(xlanguage.Japanese, "マリエハムン")
}

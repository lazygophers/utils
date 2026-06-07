//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.Japanese, "ニジェール")
	dataNiger.RegisterOfficialName(xlanguage.Japanese, "ニジェール共和国")
	dataNiger.RegisterCapital(xlanguage.Japanese, "ニアメ")
}

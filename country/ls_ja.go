//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Japanese, "レソト")
	dataLesotho.RegisterOfficialName(xlanguage.Japanese, "レソト王国")
	dataLesotho.RegisterCapital(xlanguage.Japanese, "マセル")
}

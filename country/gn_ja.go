//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Japanese, "ギニア")
	dataGuinea.RegisterOfficialName(xlanguage.Japanese, "ギニア共和国")
	dataGuinea.RegisterCapital(xlanguage.Japanese, "コナクリ")
}

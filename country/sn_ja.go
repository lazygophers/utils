//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Japanese, "セネガル")
	dataSenegal.RegisterOfficialName(xlanguage.Japanese, "セネガル共和国")
	dataSenegal.RegisterCapital(xlanguage.Japanese, "ダカール")
}

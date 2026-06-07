//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Japanese, "メキシコ")
	dataMexico.RegisterOfficialName(xlanguage.Japanese, "メキシコ合衆国")
	dataMexico.RegisterCapital(xlanguage.Japanese, "メキシコシティ")
}

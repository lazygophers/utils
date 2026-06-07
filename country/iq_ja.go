//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.Japanese, "イラク")
	dataIraq.RegisterOfficialName(xlanguage.Japanese, "イラク共和国")
	dataIraq.RegisterCapital(xlanguage.Japanese, "バグダード")
}

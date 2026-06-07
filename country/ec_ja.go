//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Japanese, "エクアドル")
	dataEcuador.RegisterOfficialName(xlanguage.Japanese, "エクアドル共和国")
	dataEcuador.RegisterCapital(xlanguage.Japanese, "キト")
}

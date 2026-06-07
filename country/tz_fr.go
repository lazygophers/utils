//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.French, "Tanzanie")
	dataTanzania.RegisterOfficialName(xlanguage.French, "République unie de Tanzanie")
	dataTanzania.RegisterCapital(xlanguage.French, "Dodoma")
}

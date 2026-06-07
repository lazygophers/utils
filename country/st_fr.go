//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.French, "Sao Tomé-et-Principe")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.French, "République démocratique de Sao Tomé-et-Principe")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.French, "Sao Tomé")
}

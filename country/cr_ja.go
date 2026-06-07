//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Japanese, "コスタリカ")
	dataCostaRica.RegisterOfficialName(xlanguage.Japanese, "コスタリカ共和国")
	dataCostaRica.RegisterCapital(xlanguage.Japanese, "サンホセ")
}

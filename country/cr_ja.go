//go:build (lang_ja || lang_all) && (country_all || country_americas || country_central_america || country_cr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Japanese, "コスタリカ")
	dataCostaRica.RegisterOfficialName(xlanguage.Japanese, "コスタリカ共和国")
	dataCostaRica.RegisterCapital(xlanguage.Japanese, "サンホセ")
}

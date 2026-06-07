//go:build country_all || country_americas || country_bl || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.English, "Saint Barthelemy")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.English, "Collectivity of Saint Barthelemy")
	dataSaintBarthelemy.RegisterCapital(xlanguage.English, "Gustavia")
}

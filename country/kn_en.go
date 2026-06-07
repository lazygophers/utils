//go:build country_all || country_americas || country_caribbean || country_kn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.English, "Saint Kitts and Nevis")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.English, "Federation of Saint Christopher and Nevis")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.English, "Basseterre")
}

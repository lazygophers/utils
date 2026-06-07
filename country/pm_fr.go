//go:build country_all || country_americas || country_northern_america || country_pm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.French, "Saint-Pierre-et-Miquelon")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.French, "Collectivité territoriale de Saint-Pierre-et-Miquelon")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.French, "Saint-Pierre")
}

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.French, "Saint-Pierre-et-Miquelon")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.French, "Collectivité territoriale de Saint-Pierre-et-Miquelon")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.French, "Saint-Pierre")
}

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.English, "Saint Pierre and Miquelon")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.English, "Overseas Collectivity of Saint Pierre and Miquelon")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.English, "Saint-Pierre")
}

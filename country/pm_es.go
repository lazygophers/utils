//go:build (lang_es || lang_all) && (country_all || country_americas || country_northern_america || country_pm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.Spanish, "San Pedro y Miquelón")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.Spanish, "Colectividad Territorial de San Pedro y Miquelón")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.Spanish, "San Pedro")
}

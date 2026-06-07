//go:build (lang_es || lang_all) && (country_all || country_by || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Spanish, "Bielorrusia")
	dataBelarus.RegisterOfficialName(xlanguage.Spanish, "República de Bielorrusia")
	dataBelarus.RegisterCapital(xlanguage.Spanish, "Minsk")
}

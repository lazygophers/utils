//go:build (lang_es || lang_all) && (country_all || country_melanesia || country_oceania || country_pg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Spanish, "Papúa Nueva Guinea")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Spanish, "Estado Independiente de Papúa Nueva Guinea")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Spanish, "Port Moresby")
}

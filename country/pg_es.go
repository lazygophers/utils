//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Spanish, "Papúa Nueva Guinea")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Spanish, "Estado Independiente de Papúa Nueva Guinea")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Spanish, "Port Moresby")
}

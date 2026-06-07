//go:build country_all || country_melanesia || country_oceania || country_pg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.English, "Papua New Guinea")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.English, "Independent State of Papua New Guinea")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.English, "Port Moresby")
}
